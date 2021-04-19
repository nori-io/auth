package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SocialProvider interface {
	Get(ctx context.Context) ([]entity.SocialProvider, error)
}
