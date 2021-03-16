package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationHistoryRepository interface {
	Create(ctx context.Context, e *entity.AuthenticationHistory) error
	Update(ctx context.Context, e *entity.AuthenticationHistory) error
	Delete(ctx context.Context, id uint64) error
}
