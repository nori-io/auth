package entity

import "time"

type MfaRecoveryCode struct {
	ID        uint64
	UserID    uint64
	Code      string
	CreatedAt time.Time
}
