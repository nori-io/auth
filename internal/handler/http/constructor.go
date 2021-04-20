package http

import (
	"github.com/nori-io/interfaces/nori/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_secret"
	"github.com/nori-plugins/authentication/internal/handler/http/settings"
	"github.com/nori-plugins/authentication/internal/handler/http/social_provider"
)

type Handler struct {
	R                      http.Http
	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
	MfaSecretHandler       *mfa_secret.MfaSecretHandler
	SettingsHandler        *settings.SettingsHandler
	SocialProviderHandler  *social_provider.SocialProviderHandler
}

type Params struct {
	R                      http.Http
	AuthenticationHandler  *authentication.AuthenticationHandler
	MfaRecoveryCodeHandler *mfa_recovery_code.MfaRecoveryCodeHandler
	MfaSecretHandler       *mfa_secret.MfaSecretHandler
	SettingsHandler        *settings.SettingsHandler
	SocialProviderHandler  *social_provider.SocialProviderHandler
}

func New(params Params) *Handler {
	handler := Handler{
		R:                      params.R,
		AuthenticationHandler:  params.AuthenticationHandler,
		MfaRecoveryCodeHandler: params.MfaRecoveryCodeHandler,
		MfaSecretHandler:       params.MfaSecretHandler,
		SettingsHandler:        params.SettingsHandler,
		SocialProviderHandler:  params.SocialProviderHandler,
	}

	// todo: add middleware
	handler.R.Post("/auth/signup", handler.AuthenticationHandler.SignUp)
	handler.R.Post("/auth/signin", handler.AuthenticationHandler.SignIn)
	handler.R.Post("/auth/signin/mfa", handler.AuthenticationHandler.SignInMfa)
	handler.R.Get("/auth/signout", handler.AuthenticationHandler.SignOut)
	handler.R.Get("/auth/session", handler.AuthenticationHandler.Session)
	handler.R.Get("/auth/settings/mfa/recovery_codes", handler.MfaRecoveryCodeHandler.GetMfaRecoveryCodes)
	handler.R.Get("/auth/social_providers", handler.SocialProviderHandler.GetSocialProviders)
	handler.R.Get("/auth/social/{social_provider}", handler.AuthenticationHandler.HandleSocialProvider)
	handler.R.Get("/auth/social/{social_provider}/callback", handler.AuthenticationHandler.HandleSocialProviderCallBack)
	// handler.R.Get("/auth/social/{social_provider}/logout", handler.AuthenticationHandler.)
	return &handler
}
