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
	SignOut(ctx context.Context, data SignOutData) error
	GetSessionData(ctx context.Context, data GetSessionData) (*entity.Session, *entity.User, error)
}

type SignUpData struct {
	Email         string
	Password      string
	TokenCaptcha  string
	ActionCaptcha string
}

func (d SignUpData) Validate() error {
	return v.Errors{
		"email":    v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

type SignInData struct {
	Email    string
	Password string
}

func (d SignInData) Validate() error {
	return v.Errors{
		"email":    v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

type SignInMfaData struct {
	SessionKey string
	Code       string
}

func (d SignInMfaData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
		"code":        v.Validate(d.Code, v.Required, v.Length(6, 6)),
	}.Filter()
}

type SignOutData struct {
	SessionKey string
}

func (d SignOutData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}

type GetSessionData struct {
	SessionKey string
}

func (d GetSessionData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}
