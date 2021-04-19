package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SocialProvider interface {
	GetAllActive(ctx context.Context) ([]entity.SocialProvider, error)
	IsSocialProviderEnabled(ctx context.Context, name string) error
}
