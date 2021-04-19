package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint"`
	Code      string    `gorm:"column:email; type: VARCHAR(15); UNIQUE"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.MfaRecoveryCode {
	return &entity.MfaRecoveryCode{
		ID:        m.ID,
		UserID:    m.UserID,
		Code:      m.Code,
		CreatedAt: m.CreatedAt,
	}
}

func newModel(e *entity.MfaRecoveryCode) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Code:      e.Code,
		CreatedAt: e.CreatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "mfa_recovery_codes"
}
