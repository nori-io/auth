package authentication

import (
	"context"
	"time"
)

type Tokens interface {
	Create(ctx context.Context, userID uint64, issuer string, length uint8, ttl time.Duration) (OneTimeToken, error)
	Use(ctx context.Context, token OneTimeToken) error
	GetByUserID(ctx context.Context, userID uint64) ([]OneTimeToken, error)
	GetByUserIdIssuer(ctx context.Context, userID uint64, issuer string) ([]OneTimeToken, error)
	GetByFilter(ctx context.Context, filter TokenFilter)
}

type OneTimeToken struct {
	ID        uint64
	UserID    uint64
	Issuer    string
	Token     []byte
	TTL       time.Duration
	CreatedAt time.Time
	UsedAt    time.Time
}

type TokenFilter struct {
	UserID    uint64
	Issuer    string
	Offset    int
	Limit     int
	CreatedAt time.Time
	UsedAt    time.Time
}
