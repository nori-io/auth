package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionService interface {
	Create(ctx context.Context, s *entity.Session) error
	Update(ctx context.Context, data *entity.Session) error
	IsActiveSessionExist(ctx context.Context, sessionKey string) (bool, error)
	GetBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error)
}

type SessionCreateData struct {
	Email    string
	Password string
}

type SessionUpdateData struct {
	Email    string
	Password string
}
