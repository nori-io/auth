package auth

import (
	"context"
	"crypto/rand"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	service2 "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	s "github.com/nori-io/interfaces/nori/session"
)

type AuthenticationService struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) service2.AuthenticationService {
	return &AuthenticationService{userRepository: userRepository}
}

func (srv AuthenticationService) SignUp(ctx context.Context, data service2.SignUpData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	var user *entity.User

	user = &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	if err := srv.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data serv.SignInData) (*entity.Session, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	var err error
	user, err = srv.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	sid, err := srv.getToken()
	if err != nil {
		return nil, err
	}
	return &entity.Session{SessionKey: sid}, nil
}

func (srv *AuthenticationService) SignOut(ctx context.Context, data *entity.Session) error {
	err := srv.session.Delete([]byte(data.SessionKey))
	return err
}

func (srv *AuthenticationService) PutSecret(
	ctx context.Context, data *serv.SecretData, session entity.Session) (
	login string, issuer string, err error) {
	if err := data.Validate(); err != nil {
		return "", "", err
	}

	var mfaSecret *entity.MfaSecret

	mfaSecret = &entity.MfaSecret{
		UserID: session.UserID,
		Secret: data.Secret,
	}

	if err := srv.mfaSecretRepository.Create(ctx, mfaSecret); err != nil {
		return "", "", err
	}

	userData, err := srv.userRepository.Get(ctx, session.UserID)
	if err != nil {
		return "", "", err
	}

	if userData.Email != "" {
		login = userData.Email
	} else {
		login = userData.PhoneCountryCode + userData.PhoneNumber
	}
	return login, srv.config.Issuer, nil
}

func (srv *AuthenticationService) getToken() ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, err
	}
	if err := srv.session.Get(sid, s.SessionActive); err != nil {
		srv.session.Save(sid, s.SessionActive, 0)
		return sid, nil
	}
	return srv.getToken()
}
