package postgres

import "time"

type MfaRecoveryCode struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserUD    uint64    `gorm:"column:user_id; type: bigint"`
	Code      string    `gorm:"column:email; type: VARCHAR(15); UNIQUE"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
}
