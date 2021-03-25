package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationLogRepository struct {
	Db *gorm.DB
}

func (r *AuthenticationLogRepository) Create(tx *gorm.DB, ctx context.Context, e *entity.AuthenticationLog) error {
	modelAuthenticationLog := NewModel(e)

	lastRecord := new(model)

	if err := tx.Create(modelAuthenticationLog).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *AuthenticationLogRepository) Update(ctx context.Context, e *entity.AuthenticationLog) error {
	model := NewModel(e)
	err := r.Db.Save(model).Error

	return err
}

func (r *AuthenticationLogRepository) FindByUserId(ctx context.Context, userId uint64) (*entity.AuthenticationLog, error) {
	var (
		out = &model{}
		e   error
	)
	e = r.Db.Where("user_id=?", userId).Last(out).Error

	return out.Convert(), e
}

func (r *AuthenticationLogRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Db.Delete(&model{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
