package pkg

import (
	"context"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-plugins/authentication/internal/domain/enum/user_status"
	"time"
)

const AuthenticationInterface meta.Interface = "nori/http/Authentication"

type Authentication interface {
	SignUp(ctx context.Context, data SignUpData) (*User, error)
	SignIn(ctx context.Context, data SignInData) (*Session, error)
	SignOut(ctx context.Context, data *Session) error
}

type SignUpData struct {
	Email    string
	Password string
}

type User struct {
	Id        uint64
	Email     string
	Password  string
	Status    user_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignInData struct {
	Email    string
	Password string
}
type Session struct {
	Id []byte
}
