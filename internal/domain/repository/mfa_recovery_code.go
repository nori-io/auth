package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Create(ctx context.Context, mfaRecoveryCode []entity.MfaRecoveryCode) error
	DeleteMfaRecoveryCode(ctx context.Context, userId uint64) error
	DeleteMfaRecoveryCodes(ctx context.Context, userId uint64) error
}
