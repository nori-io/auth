package auth

import (
	"context"

	"github.com/nori-io/authentication/internal/domain/entity"

	"github.com/nori-io/authentication/internal/domain/repository"

	"github.com/cheebo/rand"
	serv "github.com/nori-io/authentication/internal/domain/service"
	h "github.com/nori-io/interfaces/nori/http"
	s "github.com/nori-io/interfaces/nori/session"
)

type service struct {
	session s.Session
	http    h.Transport
	db      repository.UserRepository
}

func New(sessionInstance s.Session, httpInstance h.Transport, dbInstance repository.UserRepository) serv.AuthenticationService {

	return &service{
		session: sessionInstance,
		http:    httpInstance,
		db:      dbInstance,
	}

}
func (srv *service) SignUp(ctx context.Context, data serv.SignUpData) (*entity.User, error) {

	err := data.Validate()
	if err != nil {
		return nil, err
	}

	var user *entity.User
	user, err = srv.db.Create(ctx, &entity.User{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *service) SignIn(ctx context.Context, data serv.SignInData) (*entity.Session, error) {
	err := data.Validate()
	if err != nil {
		return nil, err
	}

	err = srv.db.Update(ctx, &entity.User{
		Email:    data.Email,
		Password: data.Password,
	})

	sid := rand.RandomAlphaNum(32)

	srv.session.Save([]byte(sid), s.SessionActive, 0)

	if err != nil {
		return nil, err
	}
	return &entity.Session{Id: sid}, nil

}

func (srv *service) SignOut(ctx context.Context, data *entity.Session) error {
	err := srv.session.Delete([]byte(data.Id))
	return err
}
