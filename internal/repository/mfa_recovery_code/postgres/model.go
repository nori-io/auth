package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type MfaRecoveryCode struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint"`
	Code      string    `gorm:"column:email; type: VARCHAR(15); UNIQUE"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
}

func (u *MfaRecoveryCode) Convert() *entity.MfaRecoveryCode {
	return &entity.MfaRecoveryCode{
		ID:        u.ID,
		UserID:    u.UserID,
		Code:      u.Code,
		CreatedAt: u.CreatedAt,
	}
}

func NewModel(e *entity.MfaRecoveryCode) *MfaRecoveryCode {
	return &MfaRecoveryCode{
		ID:        e.ID,
		UserID:    e.UserID,
		Code:      e.Code,
		CreatedAt: e.CreatedAt,
	}
}

// TableName
func (MfaRecoveryCode) TableName() string {
	return "mfa_recovery_codes"
}
