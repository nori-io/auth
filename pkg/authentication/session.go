package authentication

import (
	"context"
	"net/http"
	"time"
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
	Offset    int
	Limit     int
	StartedAt time.Time
	EndedAt   time.Time
}
