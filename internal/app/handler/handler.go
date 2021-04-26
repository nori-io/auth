package handler

import (
	"github.com/google/wire"
	httpHandler "github.com/nori-plugins/authentication/internal/handler/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_secret"
	"github.com/nori-plugins/authentication/internal/handler/http/settings"
	"github.com/nori-plugins/authentication/internal/handler/http/social_provider"
)

var HandlerSet = wire.NewSet(
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationHandler", "MfaRecoveryCodeHandler",
		"MfaSecretHandler", "SettingsHandler", "SocialProviderHandler", "GothProviderHelper", "SocialProviderService"),
	wire.Struct(new(authentication.Params), "AuthenticationService", "SessionService", "Logger",
		"Config", "CookieHelper", "ErrorHelper"),
	authentication.New,
	wire.Struct(new(mfa_recovery_code.Params), "MfaRecoveryCodeService", "Logger", "CookieHelper", "ErrorHelper"),
	mfa_recovery_code.New,
	wire.Struct(new(mfa_secret.Params), "MfaSecretService", "Logger", "CookieHelper", "ErrorHelper"),
	mfa_secret.New,
	wire.Struct(new(settings.Params), "SettingsService", "Logger", "CookieHelper", "ErrorHelper"),
	settings.New,
	wire.Struct(new(social_provider.Params), "SocialProviderService", "Logger", "CookieHelper", "ErrorHelper"),
	social_provider.New,
)
