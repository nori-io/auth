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
	SignInByToken(ctx context.Context, data SignInByTokenData) (*Session, error)
	SignOut(ctx context.Context, data *Session) error

	SignInSocial(w http.ResponseWriter, req http.Request) (resp *SignInSocialResponse)
	SignOutSocial(w http.ResponseWriter, req http.Request) (resp *SignOutSocialResponse)

	GetUserById(ctx context.Context, userID uint64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByPhohe(ctx context.Context, phoneNumber string) (*User, error)

	GetUserStatus(ctx context.Context, userID uint64) user_status.UserStatus


	IssueAuthenticationToken(ctx context.Context, userID uint64, length uint8, expireDuration uint16)
	GetAuthenticationToken(ctx context.Context, userID uint64)
	VerifyAuthenticationToken(ctx context.Context, data SignInByTokenData)

	GetCurrentUser(ctx context.Context)(*User, error)
	GetCurrentSession(ctx context.Context) (*Session, error)
	GetActiveSessions(ctx context.Context, userID uint64) ([]Session, error)
	OpenSession(w http.ResponseWriter, s *Session) error
	CloseActiveSessions(w http.ResponseWriter, userID uint64) error
}

type SignUpData struct {
	Login    string
	Password string
}

type User struct {
	ID        uint64
	Email     string
	Phone string
	Password  string
	Status    user_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignInByTokenData struct {
	Token    string
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