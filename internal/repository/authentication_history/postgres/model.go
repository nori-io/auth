package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SigninAt  time.Time `gorm:"column:signin_at; type: TIMESTAMP; not null"`
	Meta      string    `gorm:"column:meta; type:VARCHAR(254)"`
	SignoutAt time.Time `gorm:"column:signout_at; type: TIMESTAMP"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m model) Convert() *entity.AuthenticationHistory {
	return &entity.AuthenticationHistory{
		ID:        m.ID,
		UserID:    m.UserID,
		SigninAt:  m.SigninAt,
		Meta:      m.Meta,
		SignoutAt: m.SignoutAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func NewModel(e *entity.AuthenticationHistory) model {
	return model{
		ID:        e.ID,
		UserID:    e.UserID,
		SigninAt:  e.SigninAt,
		Meta:      e.Meta,
		SignoutAt: e.SignoutAt,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "authentication_history"
}
