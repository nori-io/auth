package postgres

import "time"

type model struct {
	ID                uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID            uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExternalID        uint64    `gorm:"column:external_id; type: bigint; not null"`
	ServiceProviderID uint8     `gorm:"column:service_provider_id; type: smallint; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProviderUserKey   string    `gorm:"column:provider_user_key; type:VARCHAR(128); not null"`
	FirstName         string    `gorm:"column:first_name; type:VARCHAR(32)"`
	LastName          string    `gorm:"column:last_name; type:VARCHAR(32)"`
	FullName          string    `gorm:"column:full_name; type:VARCHAR(64)"`
	Email             string    `gorm:"column:email; type:VARCHAR(254)"`
	AvatarURL         string    `gorm:"column:avatar_url; type:VARCHAR(254)"`
	CreatedAt         time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}
