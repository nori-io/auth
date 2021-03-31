package service

import "context"

type SessionService interface {
	IsSessionExist(ctx context.Context, sessionKey string) (bool, error)
}
