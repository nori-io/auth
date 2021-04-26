package service

import (
	"context"
)

type MfaSecretService interface {
	PutSecret(ctx context.Context, data SecretData) (
		string, string, error)
}

type SecretData struct {
	Secret     string
	SessionKey string
}

func (d SecretData) Validate() error {
	return nil
}
