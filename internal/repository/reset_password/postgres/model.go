package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token     string    `gorm:"column:token; type:VARCHAR(32); not null"`
	ExpiresAt time.Time `gorm:"column:expires_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.ResetPassword {
	return &entity.ResetPassword{
		ID:        m.ID,
		UserID:    m.UserID,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
	}
}

func newModel(e *entity.ResetPassword) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Token:     e.Token,
		ExpiresAt: e.ExpiresAt,
	}
}

// TableName
func (model) TableName() string {
	return "nori_authentication_reset_password"
}
