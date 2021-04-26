package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type SettingsService interface {
	ReceiveMfaStatus(ctx context.Context, data ReceiveMfaStatusData) (*bool, error)
	DisableMfa(ctx context.Context, data DisableMfaData) error
	ChangePassword(ctx context.Context, data ChangePasswordData) error
}

type ReceiveMfaStatusData struct {
	SessionKey string
}

func (d ReceiveMfaStatusData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}

type DisableMfaData struct {
	SessionKey string
}

func (d DisableMfaData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}

type ChangePasswordData struct {
	SessionKey  string
	PasswordOld string
	PasswordNew string
}

func (d ChangePasswordData) Validate() error {
	return v.Errors{
		"session_key":  v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
		"password_old": v.Validate(d.PasswordOld, v.Required),
		"password_new": v.Validate(d.PasswordOld, v.Required),
	}.Filter()
}
