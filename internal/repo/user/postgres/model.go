package postgres

import (
	"github.com/nori-io/authentication/internal/domain/entity"
	"time"
)


type User struct {
	Id            uint64  `gorm:"column:id; PRIMARY_KEY; type:bigserial" json:"id"`
	Email         string  `gorm:"column:email; type: VARCHAR(64)" json:"email"`
	Password	  string  `gorm:"column:email; type: VARCHAR(64)" json:"email"`
	ProfileTypeId int64   `gorm:"column:user_id; type:bigint" json:"user_id"`
	StatusId      int64	  `gorm:"column:user_id; type:bigint" json:"user_id"`
	Kind          string	`gorm:"column:status; type: VARCHAR(16)" json:"status"`
	Created       time.Time `gorm:"column:created_at; type: TIMESTAMP" json:"created_at"`
	Updated       time.Time	`gorm:"column:updated_at; type: TIMESTAMP" json:"updated_at"`
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

