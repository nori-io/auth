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

func (m MfaRecoveryCodeRepository) Create(ctx context.Context, e *entity.MfaRecoveryCode) error {
	model, _ := NewModel(e)

	lastRecord := new(MfaRecoveryCode)

	if err := m.Db.Create(model).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}
