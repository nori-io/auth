package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaSecret struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Secret    string    `gorm:"column:email; type: VARCHAR(128); UNIQUE"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m *MfaSecret) Convert() *entity.MfaSecret {
	return &entity.MfaSecret{
		ID:        m.ID,
		UserID:    m.UserID,
		Secret:    m.Secret,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func NewModel(e *entity.MfaSecret) *MfaSecret {
	return &MfaSecret{
		ID:        e.ID,
		UserID:    e.UserID,
		Secret:    e.Secret,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// Table Name
func (MfaSecret) TableName() string {
	return "mfa_secrets"
}
