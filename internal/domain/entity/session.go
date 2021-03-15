package entity

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
)

type Session struct {
	ID         uint64
	UserID     uint64
	SessionKey []byte
	Status     session_status.SessionStatus
	OpenedAt   time.Time
	ClosedAt   time.Time
	UpdatedAt  time.Time
}
