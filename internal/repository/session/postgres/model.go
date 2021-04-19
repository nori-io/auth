package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID     uint64 `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID uint64 `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//@todo возможно sessionkey сделать массивом байтов
	SessionKey string    `gorm:"column:session_key; type:VARCHAR(128); not null"`
	Status     uint8     `gorm:"column:status; type:smallint; not null"`
	OpenedAt   time.Time `gorm:"column:opened_at; type: TIMESTAMP; not null"`
	ClosedAt   time.Time `gorm:"column:closed_at; type: TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.Session {
	return &entity.Session{
		ID:         m.ID,
		UserID:     m.UserID,
		SessionKey: []byte(m.SessionKey),
		Status:     session_status.SessionStatus(m.Status),
		OpenedAt:   m.OpenedAt,
		ClosedAt:   m.ClosedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func newModel(e *entity.Session) *model {
	return &model{
		ID:         e.ID,
		UserID:     e.UserID,
		SessionKey: string(e.SessionKey),
		Status:     uint8(e.Status),
		OpenedAt:   e.OpenedAt,
		ClosedAt:   e.ClosedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "sessions"
}
