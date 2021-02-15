package entity

import "time"

type OneTimeToken struct {
	ID        uint64
	UserID    uint64
	Issuer    string
	Token     []byte
	TTL       time.Duration
	CreatedAt time.Time
	UsedAt    time.Time
}