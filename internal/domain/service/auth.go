package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	SignIn(ctx context.Context, data SignInData) (*entity.Session, error)
	SignInMfa(ctx context.Context, data SignInMfaData) (*entity.Session, uint8, error)
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
		"email":    v.Validate(d.Title, v.Required, v.Length(2, 60)),
		"password": v.Validate(d.Template, v.Required),
		//@todo проверки для каптчи?
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
