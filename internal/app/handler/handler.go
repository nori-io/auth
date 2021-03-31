package handler

import (
	"github.com/google/wire"
	httpHandler "github.com/nori-plugins/authentication/internal/handler/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_secret"
)

var HandlerSet = wire.NewSet(
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationHandler", "MfaRecoveryCodeHandler", "MfaSecretHandler"),
	wire.Struct(new(authentication.Params), "AuthenticationService", "Logger"),
	wire.Struct(new(mfa_recovery_code.Params), "MfaRecoveryCodeService", "Logger"),
	authentication.New,
	mfa_recovery_code.New,
	mfa_secret.New,
)
