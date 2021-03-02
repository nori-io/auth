package http

import (
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
)

type Handler struct {
	R                      http.Http
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	UrlPrefix              string
}

func New(h Handler) *Handler {
	handler := Handler{
		R:                      h.R,
		AuthenticationService:  h.AuthenticationService,
		MfaRecoveryCodeService: h.MfaRecoveryCodeService,
		UrlPrefix:              h.UrlPrefix,
	}
	// todo: add middleware
	Start(h)
	return &handler
}

func Start(h Handler) {
	handlerAuthentication := authentication.New(h.AuthenticationService)
	handlerMfaRecoveryCode := mfa_recovery_code.New(h.MfaRecoveryCodeService)

	h.R.Get("/auth/signup", handlerAuthentication.SignUp)
	h.R.Get("/auth/signin", handlerAuthentication.SigIn)
	h.R.Get("/auth/signout", handlerAuthentication.SignOut)

	// mfa
	h.R.Get("/auth/settings/mfa", nil)
	// h.R.Get("/auth/settings/mfa/verify?", handler.PutSecret)
	h.R.Get("/auth/settings/mfa/recovery_codes", handlerMfaRecoveryCode.GetMfaRecoveryCodes)
}
