package auth

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/nori-plugins/authentication/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	service "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *AuthenticationService) GetSessionInfo(ctx context.Context, ssid string) (*entity.Session, *entity.User, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, ssid)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	user, err := srv.userRepository.FindById(ctx, session.UserID)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}
	return &entity.Session{
			OpenedAt: session.OpenedAt,
		},
		&entity.User{
			PhoneCountryCode: user.PhoneCountryCode,
			PhoneNumber:      user.PhoneNumber,
			Email:            user.Email,
		}, nil
}

//@todo add checking action and token captcha
func (srv AuthenticationService) SignUp(ctx context.Context, data service.SignUpData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, errors.New("invalid_data", err.Error(), errors.ErrValidation)
	}

	if srv.config.EmailVerification() {
		//@todo задействовать зависимость от плагина с интерфейсом mail
	}
	tx := srv.db.Begin()

	user, err := srv.userService.CreateUser(tx, ctx, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = srv.authenticationLogService.CreateAuthenticationLog(tx, ctx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data service.SignInData) (*entity.Session, *string, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, errors.New("invalid_data", err.Error(), errors.ErrValidation)
	}

	user, err := srv.userRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	sid, err := srv.getToken(ctx)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	tx := srv.db.Begin()

	if err := srv.sessionRepository.Create(tx, ctx, &entity.Session{
		ID:         0,
		UserID:     user.ID,
		SessionKey: sid,
		Status:     session_status.Active,
		OpenedAt:   time.Now(),
	}); err != nil {
		tx.Rollback()
		return nil, nil, errors.NewInternal(err)
	}

	session, err := srv.sessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		tx.Rollback()
		return nil, nil, errors.NewInternal(err)
	}

	if err = srv.authenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
		ID:     0,
		UserID: user.ID,
		Action: users_action.SignIn,
		//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
		Meta:      "",
		SessionID: session.ID,
		CreatedAt: time.Now(),
	}); err != nil {

		tx.Rollback()
		return nil, nil, errors.NewInternal(err)
	}

	tx.Commit()

	mfaType := user.MfaType.String()

	if mfaType == mfa_type.Phone.String() {
		//@todo послать смс на номер пользователя
	}

	return &entity.Session{
		SessionKey: sid,
	}, &mfaType, nil
}

func (srv *AuthenticationService) SignInMfa(ctx context.Context, data service.SignInMfaData) (*entity.Session, error) {
	var err error
	if err = data.Validate(); err != nil {
		return nil, errors.New("invalid_data", err.Error(), errors.ErrValidation)
	}

	var session *entity.Session

	//@todo проверить кэш и отп
	isCodeFounded := srv.mfaRecoveryCodeRepository.FindByUserIdMfaRecoveryCode(ctx, session.UserID, data.Code)

	if !isCodeFounded {
		return nil, err
	}
	if isCodeFounded {
		err = srv.mfaRecoveryCodeRepository.DeleteMfaRecoveryCode(ctx, session.UserID, data.Code)
		if err != nil {
			return nil, err
		}
	}

	sid, err := srv.getToken()

	tx := srv.db.Begin()
	if err := srv.sessionRepository.Create(tx, ctx, &entity.Session{
		ID:         0,
		UserID:     session.UserID,
		SessionKey: sid,
		Status:     session_status.Active,
		OpenedAt:   time.Now(),
	}); err != nil {
		tx.Rollback()
		//@todo тут тоже возвращается ошибка, как быть?
		//создать новый тип ошибки?
		//и как быть в коде дальше
		return nil, err
	}

	session, err = srv.sessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = srv.authenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
		ID:     0,
		UserID: session.UserID,
		Action: users_action.SignInMfa,
		//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
		Meta:      "",
		SessionID: session.ID,
		CreatedAt: time.Now(),
	}); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &entity.Session{
		SessionKey: sid,
	}, nil
}

func (srv *AuthenticationService) SignOut(ctx context.Context, sess *entity.Session) error {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, string(sess.SessionKey))
	if err != nil {
		return err
	}

	tx := srv.db.Begin()
	//@todo если передать не все поля, то обнулятся ли непереданные поля в базе данных?
	if err := srv.sessionRepository.Update(tx, ctx, &entity.Session{
		ID:        session.ID,
		UserID:    session.UserID,
		Status:    session_status.Inactive,
		ClosedAt:  time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err := srv.authenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
		ID:        0,
		UserID:    session.UserID,
		Action:    users_action.SignOut,
		SessionID: session.ID,
		CreatedAt: time.Now(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (srv *AuthenticationService) getToken(ctx context.Context) ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, err
	}

	sess, err := srv.sessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		return nil, err
	}
	if sess != nil && sess.Status == session_status.Active {
		return srv.getToken(ctx)
	}
	return sid, nil
}
