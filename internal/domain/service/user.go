package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserService interface {
	Create(ctx context.Context, data UserCreateData) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, ID uint64) (*entity.User, error)
}

type UserCreateData struct {
	Email    string
	Password string
}
