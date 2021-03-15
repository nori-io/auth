package postgres

import "time"

type model struct {
	ID         uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID     uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SessionKey string    `gorm:"column:session_key; type:VARCHAR(128); not null"`
	Status     uint8     `gorm:"column:status; type:smallint; not null"`
	OpenedAt   time.Time `gorm:"column:opened_at; type: TIMESTAMP; not null"`
	ClosedAt   time.Time `gorm:"column:closed_at; type: TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}
