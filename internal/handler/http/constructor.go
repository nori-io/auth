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
	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
}

func New(h Handler, ha *authentication.AuthenticationHandler, hm *mfa_recovery_code.MfaRecoveryCodeHandler) *Handler {
	handler := Handler{
		R:                      h.R,
		AuthenticationService:  h.AuthenticationService,
		MfaRecoveryCodeService: h.MfaRecoveryCodeService,
		UrlPrefix:              h.UrlPrefix,
		AuthenticationHandler:  ha,
		MfaRecoveryCodeHandler: hm,
	}

	// todo: add middleware
	Start(h)
	return &handler
}

func Start(h Handler) {
	h.R.Get("/auth/signup", h.AuthenticationHandler.SignUp)
	h.R.Get("/auth/signin", h.AuthenticationHandler.SignIn)
	h.R.Get("/auth/signout", h.AuthenticationHandler.SignOut)

	// mfa
	h.R.Get("/auth/settings/mfa", nil)
	// h.R.Get("/auth/settings/mfa/verify?", handler.PutSecret)
	h.R.Get("/auth/settings/mfa/recovery_codes", h.MfaRecoveryCodeHandler.GetMfaRecoveryCodes)
}
