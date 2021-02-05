package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, error)
	SignOut(ctx context.Context, data *entity.Session) error
	MfaRecoveryCodes(ctx context.Context, data *entity.Session) error
}

type SignUpData struct {
	Email    string
	Password string
}

type SignInData struct {
	Email    string
	Password string
}

func (d SignUpData) Validate() error {
	return nil
}

func (d SignInData) Validate() error {
	return nil
}
