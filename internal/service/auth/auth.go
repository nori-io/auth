package auth

import (
	"context"
	"crypto/rand"
	"time"

	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"
	"github.com/nori-plugins/authentication/pkg/enum/users_action"
	"github.com/nori-plugins/authentication/pkg/enum/users_status"
	"github.com/nori-plugins/authentication/pkg/enum/users_type"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	s "github.com/nori-io/interfaces/nori/session"

	service "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *AuthenticationService) GetSessionInfo(ctx context.Context, ssid string) (*entity.Session, *entity.User, error) {
	session, err := srv.SessionRepository.FindBySessionKey(ctx, ssid)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	user, err := srv.UserRepository.FindById(ctx, session.UserID)
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

	if srv.Config.EmailVerification() {
		//@todo задействовать зависимость от плагина с интерфейсом mail
	}

	user, err := srv.UserRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	if user != nil {
		return nil, errors2.DuplicateUser
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), srv.Config.PasswordBcryptCost())

	//@todo заполнить оставшиеся поля по мере разработки нового функционала
	user = &entity.User{
		Status:          users_status.Active,
		UserType:        users_type.User,
		MfaType:         mfa_type.None,
		Email:           data.Email,
		Password:        string(password),
		HashAlgorithm:   hash_algorithm.Bcrypt,
		IsEmailVerified: srv.Config.EmailVerification(),
		CreatedAt:       time.Now(),
	}

	tx := srv.DB.Begin()
	if err := srv.UserRepository.Create(tx, ctx, user); err != nil {
		tx.Rollback()
		return nil, errors.NewInternal(err)
	}

	authenticationLog := &entity.AuthenticationLog{
		UserID: user.ID,
		Action: users_action.SignUp,
		//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
		CreatedAt: time.Now(),
	}

	if err = srv.AuthenticationLogRepository.Create(tx, ctx, authenticationLog); err != nil {
		tx.Rollback()
		return nil, errors.NewInternal(err)
	}

	tx.Commit()

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data service.SignInData) (*entity.Session, *string, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, errors.New("invalid_data", err.Error(), errors.ErrValidation)
	}

	user, err := srv.UserRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	sid, err := srv.getToken()
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	tx := srv.DB.Begin()

	if err := srv.SessionRepository.Create(tx, ctx, &entity.Session{
		ID:         0,
		UserID:     user.ID,
		SessionKey: sid,
		Status:     session_status.Active,
		OpenedAt:   time.Now(),
	}); err != nil {
		tx.Rollback()
		return nil, nil, errors.NewInternal(err)
	}

	session, err := srv.SessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		tx.Rollback()
		return nil, nil, errors.NewInternal(err)
	}

	if err = srv.AuthenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
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
	isCodeFounded := srv.MfaRecoveryCodeRepository.FindByUserIdMfaRecoveryCode(ctx, session.UserID, data.Code)

	if !isCodeFounded {
		return nil, err
	}
	if isCodeFounded {
		err = srv.MfaRecoveryCodeRepository.DeleteMfaRecoveryCode(ctx, session.UserID, data.Code)
		if err != nil {
			return nil, err
		}
	}

	sid, err := srv.getToken()

	tx := srv.DB.Begin()
	if err := srv.SessionRepository.Create(tx, ctx, &entity.Session{
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

	session, err = srv.SessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = srv.AuthenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
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
	session, err := srv.SessionRepository.FindBySessionKey(ctx, string(sess.SessionKey))
	if err != nil {
		return err
	}

	tx := srv.DB.Begin()
	//@todo если передать не все поля, то обнулятся ли непереданные поля в базе данных?
	if err := srv.SessionRepository.Update(tx, ctx, &entity.Session{
		ID:        session.ID,
		UserID:    session.UserID,
		Status:    session_status.Inactive,
		ClosedAt:  time.Now(),
		UpdatedAt: time.Now(),
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err := srv.AuthenticationLogRepository.Create(tx, ctx, &entity.AuthenticationLog{
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

func (srv *AuthenticationService) getToken() ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, err
	}
	// @todo что сохраняем в сессии?
	if err := srv.Session.Get(sid, s.SessionActive); err != nil {
		srv.Session.Save(sid, s.SessionActive, 0)
		return sid, nil
	}
	return srv.getToken()
}
