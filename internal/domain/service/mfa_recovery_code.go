package service

import (
	"context"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeService interface {
	GetMfaRecoveryCodes(ctx context.Context, data GetMfaRecoveryCodes) ([]*entity.MfaRecoveryCode, error)
	Apply(ctx context.Context, data ApplyData) error
}

type GetMfaRecoveryCodes struct {
	SessionKey string
}

func (d GetMfaRecoveryCodes) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}

type ApplyData struct {
	UserID uint64
	Code   string
}

func (d ApplyData) Validate() error {
	return v.Errors{
		"user_ID": v.Validate(d.UserID, v.Required),
		"code":    v.Validate(d.Code, v.Required),
	}.Filter()
}
