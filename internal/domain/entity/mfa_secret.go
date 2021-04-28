package entity

import "time"

type MfaSecret struct {
	ID        uint64
	UserID    uint64
	Secret    string
	CreatedAt time.Time
}
