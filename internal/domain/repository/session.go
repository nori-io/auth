package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionRepository interface {
	Create(tx *gorm.DB, ctx context.Context, e *entity.Session) error
	Update(ctx context.Context, e *entity.Session) error
	FindBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error)
}
