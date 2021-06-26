package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserLogRepository interface {
	Create(ctx context.Context, e *entity.UserLog) error
	Update(ctx context.Context, e *entity.UserLog) error
	Delete(ctx context.Context, ID uint64) error
}
