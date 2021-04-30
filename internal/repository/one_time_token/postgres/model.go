package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type model struct {
	ID     uint64 `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID uint64 `gorm:"column:user_id; type: bigint; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Issuer string `gorm:"column:issuer; type:VARCHAR(64)"`
	//@todo токен в модели - последовательность байт, проверить корректность конвертации, возможно в бд сделать массив байт
	Token     string    `gorm:"column:token; type:VARCHAR(254)"`
	TTL       time.Time `gorm:"column:ttl; type: TIMESTAMP; not null"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UsedAt    time.Time `gorm:"column:used_at; type: TIMESTAMP"`
}

func (m *model) convert() *entity.OneTimeToken {
	return &entity.OneTimeToken{
		ID:        m.ID,
		UserID:    m.UserID,
		Issuer:    m.Issuer,
		Token:     []byte(m.Token),
		TTL:       m.TTL,
		CreatedAt: m.CreatedAt,
		UsedAt:    m.UsedAt,
	}
}

func newModel(e *entity.OneTimeToken) *model {
	return &model{
		ID:        e.ID,
		UserID:    e.UserID,
		Issuer:    e.Issuer,
		Token:     string(e.Token),
		TTL:       e.TTL,
		CreatedAt: e.CreatedAt,
		UsedAt:    e.UsedAt,
	}
}

// TableName
func (model) TableName() string {
	return "nori_authentication_one_time_tokens"
}
