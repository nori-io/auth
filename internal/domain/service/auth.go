package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, error)
	SignInMfa(ctx context.Context, data SignInMfaData) (*entity.Session, uint8, error)
	SignOut(ctx context.Context, data *entity.Session) error
}

type SignUpData struct {
	Email    string
	Password string
}

type SignInData struct {
	Email    string
	Password string
}

type SignInMfaData struct {
	SessionKey string
	Code       string
}

//@todo ?
func (d SignUpData) Validate() error {
	return nil
}

//@todo ?
func (d SignInData) Validate() error {
	return nil
}

//@todo ?
func (d SignInMfaData) Validate() error {
	return nil
}
