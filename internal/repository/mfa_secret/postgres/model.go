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
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.MfaSecret {
	return &entity.MfaSecret{
		ID:        m.ID,
		UserID:    m.UserID,
		Secret:    m.Secret,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func newModel(e *entity.MfaSecret) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Secret:    e.Secret,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// Table Name
func (model) TableName() string {
	return "mfa_secrets"
}
