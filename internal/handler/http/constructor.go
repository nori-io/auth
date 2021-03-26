package http

import (
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_secret"
	"github.com/nori-plugins/authentication/internal/handler/http/settings"
)

type Handler struct {
	R                      http.Http
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	Config                 config.Config
	Logger                 logger.FieldLogger
	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
	MfaSecretHandler       *mfa_secret.MfaSecretHandler
	SettingsHandler        *settings.SettingsHandler
}

type Params struct {
	R                      http.Http
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	Config                 config.Config
	Logger                 logger.FieldLogger

	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
	MfaSecretHandler       *mfa_secret.MfaSecretHandler
	SettingsHandler        *settings.SettingsHandler
}

func New(params Params) *Handler {
	handler := Handler{
		R:                      params.R,
		AuthenticationHandler:  params.AuthenticationHandler,
		MfaRecoveryCodeHandler: params.MfaRecoveryCodeHandler,
		MfaSecretHandler:       params.MfaSecretHandler,
		SettingsHandler:        params.SettingsHandler,
	}

	// todo: add middleware
	handler.R.Get("/auth/signup", handler.AuthenticationHandler.SignUp)
	handler.R.Get("/auth/signin", handler.AuthenticationHandler.SignIn)
	handler.R.Get("/auth/signin/mfa", handler.AuthenticationHandler.SignInMfa)
	handler.R.Get("/auth/signout", handler.AuthenticationHandler.SignOut)

	// mfa
	// handler.R.Get("/auth/settings/mfa", handler.SettingsHandler)

	// h.R.Get("/auth/settings/mfa/verify?", handler.PutSecret)
	handler.R.Get("/auth/settings/mfa/recovery_codes", handler.MfaRecoveryCodeHandler.GetMfaRecoveryCodes)
	return &handler
}
