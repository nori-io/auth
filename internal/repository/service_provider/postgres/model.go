package postgres

import "time"

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Name      string    `gorm:"column:name; type:VARCHAR(64)"`
	Logo      string    `gorm:"column:logo; type:VARCHAR(254)"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}
