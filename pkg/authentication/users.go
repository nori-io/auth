package authentication

import (
	"context"
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"
)

type Users interface {
	GetUserById(ctx context.Context, userID uint64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByPhone(ctx context.Context, phone string) (User, error)
	GetCurrentUser(ctx context.Context) (User, error)
	GetUserStatus(ctx context.Context, userID uint64) users_status.UserStatus
}

type User struct {
	ID        uint64
	Email     string
	Phone     string
	Status    users_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
