package authentication

import (
	"context"
	"time"
)

type Tokens interface {
	Create(ctx context.Context, userID uint64, lengthTokenAccess uint8, lengthTokenRefresh, ttl time.Duration) (Token, error)
	Delete(ctx context.Context, token string) error
	Get(ctx context.Context, userID uint64, issuer string) (*Token, error)
	GetByUserID(ctx context.Context, userID uint64) ([]Token, error)
}

type Token struct {
	UserID    uint64
	Issuer    string
	Token     string
	TTL       time.Duration
	CreatedAt time.Time
}
