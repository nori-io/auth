package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaSecretService interface {
	PutSecret(ctx context.Context, data *SecretData, session entity.Session) (
		string, string, error)
}

type SecretData struct {
	Secret string
	Ssid   string
}

func (d SecretData) Validate() error {
	return nil
}
