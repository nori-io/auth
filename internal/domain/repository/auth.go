package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthRepository interface {
	Create(ctx context.Context, e *entity.User) error
	Get(ctx context.Context, id uint64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetAll(ctx context.Context, offset uint64, limit uint64) ([]entity.User, error)
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, id uint64) error
}
