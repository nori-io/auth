package repository

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, e *entity.User) error
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, ID uint64) error
	FindByID(ctx context.Context, ID uint64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
	FindByFilter(ctx context.Context, filter UserFilter) ([]entity.User, error)
	Count(ctx context.Context) (uint64, error)
}

type UserFilter struct {
	EmailPattern *string
	PhonePattern *string
	UserStatus   *users_status.UserStatus
	Offset       int
	Limit        int
}
