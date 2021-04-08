package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type AuthenticationLogService interface {
	Create(ctx context.Context, user *entity.User) error
}
