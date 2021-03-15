package postgres

import "time"

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID    uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Issuer    string    `gorm:"column:issuer; type:VARCHAR(64)"`
	Token     string    `gorm:"column:token; type:VARCHAR(254)"`
	TTL       time.Time `gorm:"column:ttl; type: TIMESTAMP; not null"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UsedAt    time.Time `gorm:"column:used_at; type: TIMESTAMP"`
}
