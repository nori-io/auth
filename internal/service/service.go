package service

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/authentication"
)

func (s Service) SignUp(ctx context.Context, data authentication.SignUpData) (authentication.User, error) {
	panic("implement me")
}

func (s Service) SignInByToken(ctx context.Context, token string) (authentication.Session, error) {
	panic("implement me")
}

func (s Service) Token() authentication.Tokens {
	panic("implement me")
}

func (s Service) Social() authentication.Social {
	panic("implement me")
}

func (s Service) Session() authentication.Sessions {
	panic("implement me")
}

func (s Service) User() authentication.Users {
	panic("implement me")
}
