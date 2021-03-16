package auth

import (
	"context"
	"crypto/rand"

	s "github.com/nori-io/interfaces/nori/session"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	service "github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService struct {
	UserRepository repository.UserRepository
	Session        s.Session
}

type Params struct {
	UserRepository repository.UserRepository
	Session        s.Session
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{UserRepository: params.UserRepository, Session: params.Session}
}

func (srv AuthenticationService) SignUp(ctx context.Context, data service.SignUpData) (*entity.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	var user *entity.User

	user = &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	if err := srv.UserRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *AuthenticationService) SignIn(ctx context.Context, data service.SignInData) (*entity.Session, error) {
	var err error
	if err = data.Validate(); err != nil {
		return nil, err
	}

	_, err = srv.UserRepository.FindByEmail(ctx, data.Login)
	if err != nil {
		return nil, err
	}

	_, err = srv.UserRepository.FindByPhone(ctx, data.Login)
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
	err := srv.Session.Delete([]byte(data.SessionKey))
	return err
}

func (srv *AuthenticationService) getToken() ([]byte, error) {
	sid := make([]byte, 32)

	if _, err := rand.Read(sid); err != nil {
		return nil, err
	}
	if err := srv.Session.Get(sid, s.SessionActive); err != nil {
		srv.Session.Save(sid, s.SessionActive, 0)
		return sid, nil
	}
	return srv.getToken()
}
