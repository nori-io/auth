package http

import (
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
)

type Handler struct {
	r                      http.Http
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	config                 config.Config
	authenticationHandler  *authentication.AuthenticationHandler
	mfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
}

type Params struct {
	R                      http.Http
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	Config                 config.Config
	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
}

func New(params Params) *Handler {
	handler := Handler{
		r:                      params.R,
		authenticationService:  params.AuthenticationService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		config:                 params.Config,
		authenticationHandler:  params.AuthenticationHandler,
		mfaRecoveryCodeHandler: params.MfaRecoveryCodeHandler,
	}

	// todo: add middleware
	handler.r.Get("/auth/signup", handler.authenticationHandler.SignUp)
	handler.r.Get("/auth/signin", handler.authenticationHandler.SignIn)
	handler.r.Get("/auth/signout", handler.authenticationHandler.SignOut)

	// mfa
	handler.r.Get("/auth/settings/mfa", nil)
	// h.R.Get("/auth/settings/mfa/verify?", handler.PutSecret)
	handler.r.Get("/auth/settings/mfa/recovery_codes", handler.mfaRecoveryCodeHandler.GetMfaRecoveryCodes)
	return &handler
}
