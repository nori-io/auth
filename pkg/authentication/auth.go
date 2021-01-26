package authentication

import (
	"context"
	"net/http"
	"time"

	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-plugins/authentication/pkg/enum/users_status"
)

const AuthenticationInterface meta.Interface = "nori/http/Authentication"

type (
	Authentication interface {
		SignUp(ctx context.Context, data SignUpData) (User, error)
		SignInByToken(ctx context.Context, data SignInByTokenData) (Session, error)

		Token() Tokens
		Session() Sessions
		User() Users
	}

	Users interface {
		GetUserById(ctx context.Context, userID uint64) (User, error)
		GetUserByEmail(ctx context.Context, email string) (User, error)
		GetUserByPhone(ctx context.Context, phone string) (User, error)
		GetCurrentUser(ctx context.Context) (User, error)
		GetUserStatus(ctx context.Context, userID uint64) users_status.UserStatus

	}

	Tokens interface {
		Create(ctx context.Context, userID uint64, lengthTokenAccess uint8, lengthTokenRefresh, ttl time.Duration) SignInByTokenData
		Delete(ctx context.Context, token string) error
		Verify(ctx context.Context, data SignInByTokenData) error

		IsValid(ctx context.Context, data SignInByTokenData) bool
		RefreshToken(ctx context.Context, tokenRefresh string)

		GetByUserID(ctx context.Context, userID uint64) SignInByTokenData
	}

	Sessions interface {
		Open(w http.ResponseWriter, s *Session) error
		Close(w http.ResponseWriter, s *Session) error
		CloseAll(w http.ResponseWriter, userID uint64) error

		GetCurrent(ctx context.Context) (Session, error)
		GetAllActive(ctx context.Context, userID uint64) ([]Session, error)
		GetByFilter(ctx context.Context, filter SessionFilter) ([]Session, error)
	}
)

type SessionFilter struct {
	// todo: add more filter fields
	Offset int
	Limit  int
	Beginning time.Time
	End time.Time
}

type SignUpData struct {
	Login    string
	Password string
}

type User struct {
	ID        uint64
	Email     string
	Phone     string
	Password  string
	Status    users_status.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignInByTokenData struct {
	TokenAccess string
	TokenRefresh string
}
type Session struct {
	ID     []byte
	Status int
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
