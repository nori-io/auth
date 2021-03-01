package authentication

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"
)

type Users interface {
	GetByID(ctx context.Context, userID uint64) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByPhone(ctx context.Context, phone string) (User, error)
	GetCurrent(ctx context.Context) (User, error)
	GetByFilter(ctx context.Context, filter UserFilter) ([]User, error)
}

type UserFilter struct {
	EmailPattern string
	PhonePattern string
	Status       users_status.UserStatus
	Offset       int
	Limit        int
}

type User struct {
	ID        uint64
	Email     string
	Phone     string
	Status    users_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
