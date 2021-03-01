package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaSecretRepository struct {
	Db *gorm.DB
}

func (r *MfaSecretRepository) Create(ctx context.Context, e *entity.MfaSecret) error {
	model := NewModel(e)

	lastRecord := new(MfaSecret)

	if err := r.Db.Create(model).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *MfaSecretRepository) Get(ctx context.Context, userID uint64) (*entity.MfaSecret, error) {
	var (
		out = &MfaSecret{}
		e   error
	)
	e = r.Db.Where("id=?", userID).First(out).Error

	return out.Convert(), e
}

func (r *MfaSecretRepository) Update(ctx context.Context, userID uint64, e *entity.MfaSecret) error {
	model := NewModel(e)
	err := r.Db.Save(model).Error

	return err
}

func (r *MfaSecretRepository) Delete(ctx context.Context, userID uint64) error {
	if err := r.Db.Delete(&MfaSecret{UserID: userID}).Error; err != nil {
		return err
	}
	return nil
}
