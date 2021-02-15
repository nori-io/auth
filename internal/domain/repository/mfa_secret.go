package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaSecret interface {
	Create(ctx context.Context, mfaSecret *entity.MfaSecret) error
	Get(ctx context.Context, userID uint64) (*entity.MfaSecret, error)
	Update(ctx context.Context, userID uint64, mfaSecret *entity.MfaSecret) error
	Delete(ctx context.Context, userID uint64) error
}
