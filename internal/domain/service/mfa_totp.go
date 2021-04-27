package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type MfaSecretService interface {
	GetSecret(ctx context.Context, data SecretData) (
		string, error)
}

type SecretData struct {
	SessionKey string
}

func (d SecretData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}
