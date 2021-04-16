package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/social_provider_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID          uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	AppID       string    `gorm:"column:app_id; type:VARCHAR(64); not null"`
	AppSecret   string    `gorm:"column:app_secret; type:VARCHAR(256); not null"`
	Name        string    `gorm:"column:name; type:VARCHAR(64); not null"`
	Logo        string    `gorm:"column:logo; type:VARCHAR(254)"`
	RedirectUrl string    `gorm:"column:redirect_url; type:VARCHAR(2048); not null"`
	CallBackUrl string    `gorm:"column:callback_url; type:VARCHAR(2048); not null"`
	TokenUrl    string    `gorm:"column:token_url; type:VARCHAR(2048); not null"`
	Status      uint8     `gorm:"column:status; type:smallint; not null"`
	CreatedAt   time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.SocialProvider {
	return &entity.SocialProvider{
		ID:          m.ID,
		AppID:       m.AppID,
		AppSecret:   m.AppSecret,
		Name:        m.Name,
		Logo:        m.Logo,
		RedirectUrl: m.RedirectUrl,
		CallBackUrl: m.CallBackUrl,
		TokenUrl:    m.TokenUrl,
		Status:      social_provider_status.Status(m.Status),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func newModel(e *entity.SocialProvider) *model {
	return &model{
		ID:          e.ID,
		AppID:       e.AppID,
		AppSecret:   e.AppSecret,
		Name:        e.Name,
		Logo:        e.Logo,
		RedirectUrl: e.RedirectUrl,
		CallBackUrl: e.CallBackUrl,
		TokenUrl:    e.TokenUrl,
		Status:      uint8(e.Status),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "social_providers"
}
