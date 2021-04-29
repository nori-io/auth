package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ResetPasswordService interface {
	RequestResetPasswordEmail(ctx context.Context, data RequestResetPasswordEmailData) error
	SetNewPasswordByResetPasswordEmailToken(ctx context.Context, data SetNewPasswordByResetPasswordEmailTokenData) error
}

type RequestResetPasswordEmailData struct {
	Email string
}

func (d RequestResetPasswordEmailData) Validate() error {
	return v.Errors{
		"email": v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
	}.Filter()
}

type SetNewPasswordByResetPasswordEmailTokenData struct {
	Token    string
	Password string
}

func (d SetNewPasswordByResetPasswordEmailTokenData) Validate() error {
	return v.Errors{
		"token":    v.Validate(d.Token, v.Required),
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}
