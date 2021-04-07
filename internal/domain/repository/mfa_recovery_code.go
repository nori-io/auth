package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Create(ctx context.Context, mfaRecoveryCode []entity.MfaRecoveryCode) error
	FindByUserIdMfaRecoveryCode(ctx context.Context, userId uint64, code string) bool
	DeleteMfaRecoveryCode(ctx context.Context, userId uint64, code string) error
	DeleteMfaRecoveryCodes(ctx context.Context, userId uint64) error
}
