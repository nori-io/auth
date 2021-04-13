package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Create(ctx context.Context, mfaRecoveryCode []entity.MfaRecoveryCode) error
	FindByUserID(ctx context.Context, userID uint64, code string) (*entity.MfaRecoveryCode, error)
	DeleteMfaRecoveryCode(ctx context.Context, userID uint64, code string) error
	DeleteMfaRecoveryCodes(ctx context.Context, userID uint64) error
}
