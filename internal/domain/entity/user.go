package entity

import (
	"time"

	"github.com/nori-plugins/authentication/internal/domain/enum/user_status"
)

type User struct {
	ID        uint64
	Email     string
	Phone     string
	Password  string
	Salt      string
	Status    user_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
