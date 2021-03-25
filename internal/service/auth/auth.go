package auth

import (
	"context"
	"crypto/rand"
	"time"

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

//@todo add checking action and token captcha
func (srv AuthenticationService) SignUp(ctx context.Context, data service.SignUpData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	if srv.Config.EmailVerification() {
		//@todo задействовать зависимость от плагина с интерфейсом mail
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), srv.Config.PasswordBcryptCost())
	//@todo заполнить оставшиеся поля

	user := &entity.User{
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
		return nil, err
	}

	authenticationLog := &entity.AuthenticationLog{
		UserID: user.ID,
		Action: users_action.SignUp,
		//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
		CreatedAt: time.Now(),
	}

	if err = srv.AuthenticationLogRepository.Create(tx, ctx, authenticationLog); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data service.SignInData) (*entity.Session, *string, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, err
	}

	user, err := srv.UserRepository.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, nil, err
	}

	sid, err := srv.getToken()
	if err != nil {
		return nil, nil, err
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
		srv.Session.Delete(sid)
		return nil, nil, err
	}

	session, err := srv.SessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		srv.Session.Delete(sid)

		tx.Rollback()
		return nil, nil, err
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
		srv.Session.Delete(sid)

		tx.Rollback()
		return nil, nil, err
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
		return nil, err
	}

	var session *entity.Session

	err = srv.Session.Get([]byte(data.SessionKey), s.SessionActive)
	if err != nil {
		return nil, err
	}

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
	srv.Session.Delete([]byte(data.SessionKey))
	if err != nil {
		return nil, err
	}

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
		srv.Session.Delete(sid)
		return nil, err
	}

	session, err = srv.SessionRepository.FindBySessionKey(ctx, string(sid))
	if err != nil {
		srv.Session.Delete(sid)

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
		srv.Session.Delete(sid)

		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &entity.Session{
		SessionKey: sid,
	}, nil
}

func (srv *AuthenticationService) SignOut(ctx context.Context, sess *entity.Session) error {
	if err := srv.Session.Delete(sess.SessionKey); err != nil {
		return err
	}

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
