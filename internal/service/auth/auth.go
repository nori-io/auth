package auth

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	service "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *AuthenticationService) GetSessionInfo(ctx context.Context, ssid string) (*entity.Session, *entity.User, error) {
	session, err := srv.sessionService.GetBySessionKey(ctx, ssid)
	if err != nil {
		return nil, nil, err
	}

	user, err := srv.userService.GetByID(ctx, service.GetByIdData{Id: session.UserID})
	if err != nil {
		return nil, nil, err
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
		return nil, err
	}

	if srv.config.EmailVerification() {
		//@todo задействовать зависимость от плагина с интерфейсом mail
	}

	userCreateData := service.UserCreateData{
		Email:    data.Email,
		Password: data.Password,
	}

	var user *entity.User
	var err error

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		user, err = srv.userService.Create(tx, userCreateData)
		if err != nil {
			return err
		}

		err = srv.authenticationLogService.Create(tx, &entity.AuthenticationLog{
			UserID: user.ID,
			Action: users_action.SignUp,
			//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
			Meta:      "",
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data service.SignInData) (*entity.Session, *string, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, err
	}

	user, err := srv.userService.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, err
	}

	if err := srv.securityHelper.ComparePassword(user.Password, data.Password); err != nil {
		return nil, nil, err
	}

	sid, err := srv.getToken(ctx)
	if err != nil {
		return nil, nil, errors.NewInternal(err)
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.sessionService.Create(ctx, &entity.Session{
			ID:         0,
			UserID:     user.ID,
			SessionKey: sid,
			Status:     session_status.Active,
			OpenedAt:   time.Now(),
		}); err != nil {
			errors.NewInternal(err)
		}

		session, err := srv.sessionService.GetBySessionKey(ctx, string(sid))
		if err != nil {
			return errors.NewInternal(err)
		}

		if err = srv.authenticationLogService.Create(ctx, &entity.AuthenticationLog{
			ID:     0,
			UserID: user.ID,
			Action: users_action.SignIn,
			//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
			Meta:      "",
			SessionID: session.ID,
			CreatedAt: time.Now(),
		}); err != nil {
			return errors.NewInternal(err)
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

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

	//@todo проверить кэш и отп
	mfaRecoveryCode, err := srv.mfaRecoveryCodeService.GetByUserId(ctx, session.UserID, data.Code)
	if err != nil {
		return nil, err
	}

	if mfaRecoveryCode != nil {
		err = srv.mfaRecoveryCodeService.Apply(ctx, session.UserID, data.Code)
		if err != nil {
			return nil, err
		}
	}

	sid, err := srv.getToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.sessionService.Create(ctx, &entity.Session{
			ID:         0,
			UserID:     session.UserID,
			SessionKey: sid,
			Status:     session_status.Active,
			OpenedAt:   time.Now(),
		}); err != nil {
			//@todo тут тоже возвращается ошибка, как быть?
			//создать новый тип ошибки?
			//и как быть в коде дальше
			return err
		}

		session, err = srv.sessionService.GetBySessionKey(ctx, string(sid))
		if err != nil {
			return err
		}

		if err = srv.authenticationLogService.Create(ctx, &entity.AuthenticationLog{
			ID:     0,
			UserID: session.UserID,
			Action: users_action.SignInMfa,
			//@todo заполнить метаданные айпи адресом и городом или чем-то ещё?
			Meta:      "",
			SessionID: session.ID,
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.Session{
		SessionKey: sid,
	}, nil
}

func (srv *AuthenticationService) SignOut(ctx context.Context, sess *entity.Session) error {
	session, err := srv.sessionService.GetBySessionKey(ctx, string(sess.SessionKey))
	if err != nil {
		return err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		//@todo если передать не все поля, то обнулятся ли непереданные поля в базе данных?
		if err := srv.sessionService.Update(ctx, &entity.Session{
			ID:        session.ID,
			UserID:    session.UserID,
			Status:    session_status.Inactive,
			ClosedAt:  time.Now(),
			UpdatedAt: time.Now(),
		}); err != nil {
			return err
		}

		if err := srv.authenticationLogService.Create(ctx, &entity.AuthenticationLog{
			ID:        0,
			UserID:    session.UserID,
			Action:    users_action.SignOut,
			SessionID: session.ID,
			CreatedAt: time.Now(),
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (srv *AuthenticationService) getToken(ctx context.Context) ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, errors.NewInternal(err)
	}

	sess, err := srv.sessionService.GetBySessionKey(ctx, string(sid))
	if err != nil {
		return nil, err
	}
	if sess != nil && sess.Status == session_status.Active {
		return srv.getToken(ctx)
	}
	return sid, nil
}
