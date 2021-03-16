package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationHistoryRepository struct {
	Db *gorm.DB
}

func (r *AuthenticationHistoryRepository) Create(ctx context.Context, e *entity.AuthenticationHistory) error {
	modelAuthenticationHistory := NewModel(e)

	lastRecord := new(model)

	if err := r.Db.Create(modelAuthenticationHistory).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *AuthenticationHistoryRepository) Update(ctx context.Context, e *entity.AuthenticationHistory) error {
	model := NewModel(e)
	err := r.Db.Save(model).Error

	return err
}

func (r *AuthenticationHistoryRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Db.Delete(&model{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
