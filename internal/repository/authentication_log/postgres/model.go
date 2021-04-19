package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE; not null"`
	Action    uint8     `gorm:"column:action; type: smallint; not null"`
	SessionID uint64    `gorm:"column:session_id; type:bigint"`
	Meta      string    `gorm:"column:meta; type:VARCHAR(254)"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
}

func (m *model) convert() *entity.AuthenticationLog {
	return &entity.AuthenticationLog{
		ID:        m.ID,
		UserID:    m.UserID,
		Action:    users_action.Action(m.Action),
		SessionID: m.SessionID,
		Meta:      m.Meta,
		CreatedAt: m.CreatedAt,
	}
}

func newModel(e *entity.AuthenticationLog) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Action:    uint8(e.Action),
		SessionID: e.SessionID,
		Meta:      e.Meta,
		CreatedAt: e.CreatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "authentication_log"
}
