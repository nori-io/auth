package postgres

import "time"

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SigninAt  time.Time `gorm:"column:signin_at; type: TIMESTAMP; not null"`
	Meta      string    `gorm:"column:meta; type:VARCHAR(254)"`
	SignoutAt time.Time `gorm:"column:signout_at; type: TIMESTAMP"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}
