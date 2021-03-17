package postgres

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionRepository struct {
	Db *gorm.DB
}

func (r *SessionRepository) Create(ctx context.Context, e *entity.Session) error {
	modelSession := NewModel(e)

	lastRecord := new(model)

	if err := r.Db.Create(modelSession).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}
