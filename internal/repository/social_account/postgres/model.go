package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID                uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID            uint64    `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExternalID        string    `gorm:"column:external_id; type: VARCHAR(32); not null"`
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

func (m model) convert() *entity.SocialAccount {
	return &entity.SocialAccount{
		ID:                m.ID,
		User_ID:           m.UserID,
		ExternalID:        m.ExternalID,
		ServiceProviderID: m.ServiceProviderID,
		ProviderUserKey:   m.ProviderUserKey,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		FullName:          m.FullName,
		Email:             m.Email,
		AvatarURL:         m.AvatarURL,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}

func newModel(e *entity.SocialAccount) model {
	return model{
		ID:                e.ID,
		UserID:            e.User_ID,
		ExternalID:        e.ExternalID,
		ServiceProviderID: e.ServiceProviderID,
		ProviderUserKey:   e.ProviderUserKey,
		FirstName:         e.FirstName,
		LastName:          e.LastName,
		FullName:          e.FullName,
		Email:             e.Email,
		AvatarURL:         e.AvatarURL,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "social_accounts"
}
