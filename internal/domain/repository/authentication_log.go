package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationLogRepository interface {
	Create(ctx context.Context, e *entity.AuthenticationLog) error
	Update(ctx context.Context, e *entity.AuthenticationLog) error
	Delete(ctx context.Context, id uint64) error
}
