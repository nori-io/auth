package auth

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"time"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	service "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv *AuthenticationService) GetSessionData(ctx context.Context, data service.GetSessionData) (*entity.Session, *entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, err
	}

	session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
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

		err = srv.userLogService.Create(tx, service.UserLogCreateData{
			UserID:    user.ID,
			Action:    users_action.SignUp,
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

func (srv *AuthenticationService) LogIn(ctx context.Context, data service.LogInData) (*entity.Session, *string, error) {
	if err := data.Validate(); err != nil {
		return nil, nil, err
	}

	user, err := srv.userService.GetByEmail(ctx, service.GetByEmailData{Email: data.Email})
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
		if err := srv.sessionService.Create(ctx, service.SessionCreateData{
			UserID:     0,
			SessionKey: "",
			Status:     0,
			OpenedAt:   time.Time{},
		}); err != nil {
			errors.NewInternal(err)
		}

		session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: string(sid)})
		if err != nil {
			return errors.NewInternal(err)
		}

		if err = srv.userLogService.Create(ctx, service.UserLogCreateData{
			UserID:    user.ID,
			Action:    users_action.LogIn,
			SessionID: session.ID,
			Meta:      "",
			CreatedAt: time.Time{},
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

func (srv *AuthenticationService) LogInMfa(ctx context.Context, data service.LogInMfaData) (*entity.Session, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil {
		return nil, err
	}

	isValid, err := srv.mfaTotpService.Validate(ctx, service.MfaTotpValidateData{
		UserID:   session.UserID,
		PassCode: data.Code,
	})
	if err != nil {
		return nil, err
	}

	if isValid != true {
		if err = srv.mfaRecoveryCodeService.Apply(ctx, service.ApplyData{
			UserID:     session.UserID,
			Code:       data.Code,
			SessionKey: data.SessionKey,
		}); err != nil {
			return nil, err
		}
	}
	sid, err := srv.getToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.sessionService.Create(ctx, service.SessionCreateData{
			UserID:     session.UserID,
			SessionKey: string(sid),
			Status:     session_status.Active,
			OpenedAt:   time.Now(),
		}); err != nil {
			//@todo тут тоже возвращается ошибка, как быть?
			//создать новый тип ошибки?
			//и как быть в коде дальше
			return err
		}

		session, err = srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: string(sid)})
		if err != nil {
			return err
		}

		if err = srv.userLogService.Create(ctx, service.UserLogCreateData{
			UserID:    session.UserID,
			Action:    users_action.LogInMfa,
			SessionID: session.ID,
			Meta:      "",
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

func (srv *AuthenticationService) LogOut(ctx context.Context, data service.LogOutData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	session, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil {
		return err
	}

	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		//@todo если передать не все поля, то обнулятся ли непереданные поля в базе данных?
		if err := srv.sessionService.Update(ctx, service.SessionUpdateData{
			UserID:     session.UserID,
			SessionKey: data.SessionKey,
			Status:     session_status.Inactive,
			ClosedAt:   time.Now(),
			UpdatedAt:  time.Now(),
		}); err != nil {
			return err
		}

		if err := srv.userLogService.Create(ctx, service.UserLogCreateData{
			UserID:    session.UserID,
			Action:    users_action.LogOut,
			SessionID: session.ID,
			Meta:      "",
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

func (srv *AuthenticationService) IsAuthenticated(r *http.Request) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := srv.cookieHelper.GetSessionID(r)
			if err != nil {
				http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
				return
			}
			_, err = srv.sessionService.GetBySessionKey(r.Context(), service.GetBySessionKeyData{SessionKey: sid})
			if err != nil {
				http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
func (srv *AuthenticationService) IsBasicAuthenticated(realm string, creds map[string]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w, realm)
				return
			}

			credPass, credUserOk := creds[user]
			if !credUserOk || subtle.ConstantTimeCompare([]byte(pass), []byte(credPass)) != 1 {
				basicAuthFailed(w, realm)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}

func (srv *AuthenticationService) getToken(ctx context.Context) ([]byte, error) {
	token, err := srv.securityHelper.GenerateToken(32)

	sess, err := srv.sessionService.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: token})
	if err != nil {
		return nil, err
	}
	if sess != nil && sess.Status == session_status.Active {
		return srv.getToken(ctx)
	}
	return []byte(token), nil
}
