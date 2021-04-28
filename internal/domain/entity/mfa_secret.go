package entity

import "time"

type MfaTotp struct {
	ID        uint64
	UserID    uint64
	Secret    string
	CreatedAt time.Time
}
