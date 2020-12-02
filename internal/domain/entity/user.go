package entity

import (
	"time"

	"github.com/nori-io/authentication/internal/domain/enum/user_status"
)

type User struct {
	Id        uint64
	Email     string
	Password  string
	Status    user_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
