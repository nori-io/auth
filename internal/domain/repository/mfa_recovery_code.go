package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository interface {
	Create(tx *gorm.DB, ctx context.Context, mfaRecoveryCode []entity.MfaRecoveryCode) error
	FindByUserIdMfaRecoveryCode(ctx context.Context, userId uint64, code string) bool
	DeleteMfaRecoveryCode(ctx context.Context, userId uint64, code string) error
	DeleteMfaRecoveryCodes(tx *gorm.DB, ctx context.Context, userId uint64) error
}
