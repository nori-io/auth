package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, *string, error)
	SignInMfa(ctx context.Context, data SignInMfaData) (*entity.Session, error)
	SignOut(ctx context.Context, data *entity.Session) error
}

type SignUpData struct {
	Email         string
	Password      string
	TokenCaptcha  string
	ActionCaptcha string
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
	return v.Errors{
		"email":    v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

//@todo ?
func (d SignInData) Validate() error {
	return nil
}

//@todo ?
func (d SignInMfaData) Validate() error {
	return nil
}
