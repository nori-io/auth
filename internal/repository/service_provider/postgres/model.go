package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Name      string    `gorm:"column:name; type:VARCHAR(64)"`
	Logo      string    `gorm:"column:logo; type:VARCHAR(254)"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (m model) Convert() *entity.ServiceProvider {
	return &entity.ServiceProvider{
		ID:        m.ID,
		Name:      m.Name,
		Logo:      m.Logo,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func NewModel(e *entity.ServiceProvider) model {
	return model{
		ID:        e.ID,
		Name:      e.Name,
		Logo:      e.Logo,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// TableName
func (model) TableName() string {
	return "service_providers"
}
