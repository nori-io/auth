package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, error)
	SignOut(ctx context.Context, data *entity.Session) error
	GetMfaRecoveryCodes(ctx context.Context, data *entity.Session) error
	GetSecret(ctx context.Context, data *entity.Session)
}

type SignUpData struct {
	Email    string
	Password string
}

type SignInData struct {
	Email    string
	Password string
}

type SecretData struct {
	Secret string
}

func (d SignUpData) Validate() error {
	return nil
}

func (d SignInData) Validate() error {
	return nil
}

func (d SecretData) Validate() error {
	return nil
}
