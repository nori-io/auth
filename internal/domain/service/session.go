package service

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionService interface {
	IsSessionExist(ctx context.Context, sessionKey string) (bool, error)
	GetBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error)
	Create(ctx context.Context, s *entity.Session) error
}

type SessionCreateData struct {
	Email    string
	Password string
}
