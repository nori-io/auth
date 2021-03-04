package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Use(ctx context.Context, e *entity.MfaRecoveryCode) error
	Create(ctx context.Context, userID uint64, mfaRecoveryCode []entity.MfaRecoveryCode) error
}
