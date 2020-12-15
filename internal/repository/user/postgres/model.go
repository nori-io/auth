package postgres

import (
	"time"

	"github.com/nori-io/authentication/internal/domain/enum/user_status"

	"github.com/nori-io/authentication/internal/domain/entity"
)

type User struct {
	Id        uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Email     string    `gorm:"column:email; type: VARCHAR(32)"`
	Password  string    `gorm:"column:email; type: VARCHAR(32)"`
	Status    uint8     `gorm:"column:user_id; type:bigint"`
	CreatedAt time.Time `gorm:"column:created_at; type: TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (u *User) Convert() *entity.User {
	return &entity.User{
		Id:        u.Id,
		Email:     u.Email,
		Password:  u.Password,
		Status:    user_status.UserStatus(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewModel(e *entity.User) (*User, error) {
	return &User{
		Id:        e.Id,
		Email:     e.Email,
		Password:  e.Password,
		Status:    uint8(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}, nil
}

// TableName
func (User) TableName() string {
	return "users"
}
