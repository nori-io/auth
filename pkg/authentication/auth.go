package authentication

import (
	"context"

	"github.com/nori-io/common/v4/pkg/domain/meta"
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
