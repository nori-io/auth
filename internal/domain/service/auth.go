package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data SignUpData) (*entity.User, error)
	LogIn(ctx context.Context, data LogInData) (*entity.Session, *string, error)
	LogInMfa(ctx context.Context, data LogInMfaData) (*entity.Session, error)
	LogOut(ctx context.Context, data LogOutData) error
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

type LogInData struct {
	Email    string
	Password string
}

func (d LogInData) Validate() error {
	return v.Errors{
		"email":    v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}

type LogInMfaData struct {
	SessionKey string
	Code       string
}

func (d LogInMfaData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
		"code":        v.Validate(d.Code, v.Required, v.Length(6, 6)),
	}.Filter()
}

type LogOutData struct {
	SessionKey string
}

func (d LogOutData) Validate() error {
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
