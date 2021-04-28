package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type MfaTotpService interface {
	GetUrl(ctx context.Context, data MfaGetUrlData) (string, error)
	Validate(ctx context.Context, data MfaTotpValidateData) (bool, error)
}

type MfaGetUrlData struct {
	SessionKey string
}

func (d MfaGetUrlData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}

type MfaTotpValidateData struct {
	UserID   uint64
	PassCode string
}

func (d MfaTotpValidateData) Validate() error {
	return v.Errors{
		"user_ID":  v.Validate(d.UserID, v.Required),
		"passcode": v.Validate(d.PassCode, v.Required, v.Length(6, 6)),
	}.Filter()
}
