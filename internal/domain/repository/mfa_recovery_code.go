package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Use(ctx context.Context, e *entity.MfaRecoveryCode) error
	Get(ctx context.Context, userID uint64) ([]entity.MfaRecoveryCode, error)
}
