package pkg

import (
	"context"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-plugins/authentication/internal/domain/enum/user_status"
	"net/http"
	"time"
)

const AuthenticationInterface meta.Interface = "nori/http/Authentication"

type Authentication interface {
	SignUp(ctx context.Context, data SignUpData) (*User, error)
	SignIn(ctx context.Context, data SignInData) (*Session, error)
	SignOut(ctx context.Context, data *Session) error

	SignInSocial(w http.ResponseWriter, req http.Request) (resp *SignInSocialResponse)
	SignOutSocial(w http.ResponseWriter, req http.Request) (resp *SignOutSocialResponse)

	// Может везде оперировать именно user Id? на входе
	GetCurrentUser(w http.ResponseWriter, id uint)(*User, error)
	GetCurrentSessionId(w http.ResponseWriter, id uint) (*Session, error)
	GetActiveSessions(w http.ResponseWriter, id uint) (*[]Session, error)
	CreateSession(w http.ResponseWriter, id uint)(*Session, error)
	CloseActiveSessions(w http.ResponseWriter, id uint) error
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

type SignInSocialResponse struct {
	Id             uint64
	Token          string
	User           UserResponse
	MFA            string
	HttpStatusCode int
	Err            error
}

type UserResponse struct {
	UserName string
}

type SignOutSocialResponse struct {
	HttpStatusCode int
	Err            error
}