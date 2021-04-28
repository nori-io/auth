package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type MfaTotpService interface {
	GetUrl(ctx context.Context, data MfaTotpData) (string, error)
}

type MfaTotpData struct {
	SessionKey string
}

func (d MfaTotpData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}
