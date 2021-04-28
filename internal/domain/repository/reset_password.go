package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type ResetPasswordRepository interface {
	Create(ctx context.Context, resetPassword *entity.ResetPassword) error
	Delete(ctx context.Context, userID uint64) error
}
