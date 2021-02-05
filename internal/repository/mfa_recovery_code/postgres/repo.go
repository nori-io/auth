package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository struct {
	Db *gorm.DB
}

func (m MfaRecoveryCodeRepository) Use(ctx context.Context, e *entity.MfaRecoveryCode) error {
	panic("implement me")
}

func (m MfaRecoveryCodeRepository) Get(ctx context.Context, userID uint64) ([]entity.MfaRecoveryCode, error) {
	panic("implement me")
}
