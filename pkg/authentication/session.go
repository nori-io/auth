package authentication

import (
	"context"
	"net/http"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
)

type Sessions interface {
	Open(w http.ResponseWriter, s *Session) error
	Close(w http.ResponseWriter, s *Session) error
	CloseAll(w http.ResponseWriter, userID uint64) error

	GetCurrent(ctx context.Context) (Session, error)
	GetAllActive(ctx context.Context, userID uint64) ([]Session, error)
	GetByFilter(ctx context.Context, filter SessionFilter) ([]Session, error)
}

type SessionFilter struct {
	UserID   uint64
	Offset   int
	Limit    int
	Status   session_status.SessionStatus
	OpenedAt time.Time
	ClosedAt time.Time
}

type Session struct {
	ID       uint64
	Key      []byte
	UserID   uint64
	Status   session_status.SessionStatus
	OpenedAt time.Time
	ClosedAt time.Time
}
