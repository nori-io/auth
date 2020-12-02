package auth

import (
	"context"

	"github.com/nori-io/authentication/internal/domain/entity"

	"github.com/nori-io/authentication/internal/domain/repository"

	"github.com/cheebo/rand"
	serv "github.com/nori-io/authentication/internal/domain/service"
	s "github.com/nori-io/interfaces/nori/session"
)

type service struct {
	session s.Session
	db      repository.UserRepository
}

func New(sessionInstance s.Session, dbInstance repository.UserRepository) serv.AuthenticationService {

	return &service{
		session: sessionInstance,
		db:      dbInstance,
	}

}
func (srv *service) SignUp(ctx context.Context, data serv.SignUpData) (*entity.User, error) {

	err := data.Validate()
	if err != nil {
		return nil, err
	}

	var user *entity.User

	user = &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	err = srv.db.Create(ctx, user)

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

	user := &entity.User{
		Email:    data.Email,
		Password: data.Password,
	}

	err = srv.db.Update(ctx, user)

	sid, err := srv.getToken()

	if err != nil {
		return nil, err
	}
	return &entity.Session{Id: sid}, nil

}

func (srv *service) SignOut(ctx context.Context, data *entity.Session) error {
	err := srv.session.Delete([]byte(data.Id))
	return err
}

func (srv *service) getToken() (string, error) {

	sid := rand.RandomAlphaNum(32)
	err := srv.session.Get([]byte(sid), s.SessionActive)
	if err != nil {
		srv.session.Save([]byte(sid), s.SessionActive, 0)
		return sid, nil
	}
	return "", nil
}
