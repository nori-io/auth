package repository

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionRepository interface {
	Create(ctx context.Context, e *entity.Session) error
}
