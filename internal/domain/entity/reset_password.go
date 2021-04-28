package entity

import "time"

type ResetPassword struct {
	ID        uint64
	UserID    uint64
	Token     string
	ExpiresAt time.Time
}
