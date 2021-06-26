package authentication

import (
	"context"
	"net/http"

	"github.com/nori-io/common/v5/pkg/domain/meta"
)

const AuthenticationInterface meta.Interface = "nori/http/Authentication"

type (
	Authentication interface {
		SignUp(ctx context.Context, data SignUpData) (User, error)
		SignInByToken(ctx context.Context, token string) (Session, error)
		IsAuthenticated() func(next http.Handler) http.Handler
		IsBasicAuthenticated(realm string, creds map[string]string) func(next http.Handler) http.Handler
		Token() Tokens
		Social() Social
		Session() Sessions
		User() Users
	}
)

type SignUpData struct {
	Login    string
	Password string
}
