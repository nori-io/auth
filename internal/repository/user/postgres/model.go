package postgres

import (
	"time"

	"github.com/nori-io/authentication/internal/domain/entity"
)

type User struct {
	Id            uint64    `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	Email         string    `gorm:"column:email; type: VARCHAR(64)"`
	Password      string    `gorm:"column:email; type: VARCHAR(64)"`
	ProfileTypeId int64     `gorm:"column:user_id; type:bigint"`
	StatusId      int64     `gorm:"column:user_id; type:bigint"`
	Kind          string    `gorm:"column:status; type: VARCHAR(16)"`
	Created       time.Time `gorm:"column:created_at; type: TIMESTAMP"`
	Updated       time.Time `gorm:"column:updated_at; type: TIMESTAMP"`
}

func (u *User) Convert() *entity.User {
	return &entity.User{
		Id:            u.Id,
		Email:         u.Email,
		Password:      u.Password,
		ProfileTypeId: u.ProfileTypeId,
		StatusId:      u.StatusId,
		Kind:          u.Kind,
		Created:       u.Created,
		Updated:       u.Updated,
	}
}

func NewModel(e *entity.User) (*User, error) {
	return &User{
		Id:            e.Id,
		Email:         e.Email,
		Password:      e.Password,
		ProfileTypeId: e.ProfileTypeId,
		StatusId:      e.StatusId,
		Kind:          e.Kind,
		Created:       e.Created,
		Updated:       e.Updated,
	}, nil
}

// TableName
func (User) TableName() string {
	return "users"
}
