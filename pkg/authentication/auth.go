package authentication

import (
	"context"
	"time"

	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-plugins/authentication/pkg/enum/session_status"
)

const AuthenticationInterface meta.Interface = "nori/http/Authentication"

type (
	Authentication interface {
		SignUp(ctx context.Context, data SignUpData) (User, error)
		SignInByToken(ctx context.Context, token string) (Session, error)

		Token() Tokens
		Session() Sessions
		User() Users
	}
)

type SignUpData struct {
	Login    string
	Password string
}

type Session struct {
	ID        []byte
	Status    session_status.SessionStatus
	StartedAt time.Time
	EndedAt   time.Time
}
