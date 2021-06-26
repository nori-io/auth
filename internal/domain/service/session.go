package service

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionService interface {
	Create(ctx context.Context, data SessionCreateData) error
	Update(ctx context.Context, data SessionUpdateData) error
	GetBySessionKey(ctx context.Context, data GetBySessionKeyData) (*entity.Session, error)
}

type SessionCreateData struct {
	UserID     uint64
	SessionKey string
	Status     session_status.SessionStatus
	OpenedAt   time.Time
}

func (d SessionCreateData) Validate() error {
	return v.Errors{
		"user_ID":     v.Validate(d.UserID, v.Required),
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
		"status":      v.Validate(d.Status, v.Required),
	}.Filter()
}

type SessionUpdateData struct {
	UserID     uint64
	SessionKey string
	Status     session_status.SessionStatus
	ClosedAt   time.Time
	UpdatedAt  time.Time
}

func (d SessionUpdateData) Validate() error {
	return v.Errors{
		"user_ID":     v.Validate(d.UserID, v.Required),
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
		"status":      v.Validate(d.Status, v.Required),
	}.Filter()
}

type GetBySessionKeyData struct {
	SessionKey string
}

func (d GetBySessionKeyData) Validate() error {
	return v.Errors{
		"session_key": v.Validate(d.SessionKey, v.Required, v.Length(128, 128)),
	}.Filter()
}
