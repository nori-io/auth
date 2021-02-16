package postgres

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type User struct {
	ID        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Status    uint8     `gorm:"column:status; type:smallint; not null" `
	UserType  uint8     `gorm:"column:user_type; type:smallint; not null"`
	MfaType   uint8     `gorm:"column:mfa_type; type:smallint; null"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP; not null"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (u *User) Convert() *entity.User {
	return &entity.User{
		ID:        u.Id,
		Status:    users_status.UserStatus(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewModel(e *entity.User) *User {
	return &User{
		Id:        e.ID,
		Email:     e.Email,
		Password:  e.Password,
		Status:    uint8(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// TableName
func (User) TableName() string {
	return "users"
}
