package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ResetPasswordService interface {
	RequestResetPasswordEmail(ctx context.Context, data RequestResetPasswordEmailData) error
	SetNewPasswordByRestorePasswordEmailToken(ctx context.Context, password string) error
}

type RequestResetPasswordEmailData struct {
	Email string
}

func (d RequestResetPasswordEmailData) Validate() error {
	return v.Errors{
		"email": v.Validate(d.Email, v.Required, v.Length(3, 254), is.Email),
	}.Filter()
}

type SetNewPasswordByRestorePasswordEmailTokenData struct {
	Password string
}

func (d SetNewPasswordByRestorePasswordEmailTokenData) Validate() error {
	return v.Errors{
		"password": v.Validate(d.Password, v.Required),
	}.Filter()
}
