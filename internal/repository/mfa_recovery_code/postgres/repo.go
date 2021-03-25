package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCodeRepository struct {
	Db *gorm.DB
}

func (r MfaRecoveryCodeRepository) Create(tx *gorm.DB, ctx context.Context, e []entity.MfaRecoveryCode) error {
	var mfaRecoveryCodes []model

	for _, v := range e {
		mfaRecoveryCodes = append(mfaRecoveryCodes, NewModel(&v))
	}

	lastRecord := new(model)

	if err := r.Db.Create(mfaRecoveryCodes).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r MfaRecoveryCodeRepository) FindByUserIdMfaRecoveryCode(ctx context.Context, userId uint64, code string) bool {
	out := &model{}

	rows := r.Db.Where("user_id=?, code=?", userId, code).First(out).RowsAffected

	if rows == 1 {
		return true
	}

	return false
}

func (r MfaRecoveryCodeRepository) DeleteMfaRecoveryCode(ctx context.Context, userId uint64, code string) error {
	if err := r.Db.Delete(&model{UserID: userId, Code: code}).Error; err != nil {
		return err
	}
	return nil
}

func (r MfaRecoveryCodeRepository) DeleteMfaRecoveryCodes(tx *gorm.DB, ctx context.Context, userId uint64) error {
	if err := r.Db.Delete(&model{UserID: userId}).Error; err != nil {
		return err
	}
	return nil
}
