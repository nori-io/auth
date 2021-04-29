package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Secret    string    `gorm:"column:email; type: VARCHAR(128); UNIQUE"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.MfaTotp {
	return &entity.MfaTotp{
		ID:        m.ID,
		UserID:    m.UserID,
		Secret:    m.Secret,
		CreatedAt: m.CreatedAt,
	}
}

func newModel(e *entity.MfaTotp) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Secret:    e.Secret,
		CreatedAt: e.CreatedAt,
	}
}

// TableName for gorm
func (model) TableName() string {
	return "mfa_totp"
}
