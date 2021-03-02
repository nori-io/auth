package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SecretService interface {
	PutSecret(ctx context.Context, data *entity.Session) (string, string, error)
}

type SecretData struct {
	Secret string
	Ssid   string
}

func (d SecretData) Validate() error {
	return nil
}
