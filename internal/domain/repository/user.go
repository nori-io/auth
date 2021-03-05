package repository

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, e *entity.User) error
	FindById(ctx context.Context, id uint64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByPhone(ctx context.Context, email string) (*entity.User, error)
	FindByFilter(ctx context.Context, filter UserFilter) ([]entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, id uint64) error
}

type UserFilter struct {
	EmailPattern string
	PhonePattern string
	Status       users_status.UserStatus
	Offset       int
	Limit        int
}

// GetByID(ctx context.Context, userID uint64) (User, error)
// GetByEmail(ctx context.Context, email string) (User, error)
// GetByPhone(ctx context.Context, phone string) (User, error)
// GetCurrent(ctx context.Context) (User, error)
// GetByFilter(ctx context.Context, filter UserFilter) ([]User, error)
