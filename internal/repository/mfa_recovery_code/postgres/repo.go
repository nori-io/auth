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

func (m MfaRecoveryCodeRepository) Create(ctx context.Context, userID uint64, e []entity.MfaRecoveryCode) error {
	var mfaRecoveryCodes []*MfaRecoveryCode

	for _, v := range e {
		mfaRecoveryCodes = append(mfaRecoveryCodes, NewModel(&v))
	}

	lastRecord := new(MfaRecoveryCode)

	if err := m.Db.Create(mfaRecoveryCodes).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}
