package service

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/service"
	"github.com/nori-plugins/authentication/internal/service/auth"
	"github.com/nori-plugins/authentication/internal/service/authentication_log"
	"github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/service/mfa_secret"
	"github.com/nori-plugins/authentication/internal/service/session"
	"github.com/nori-plugins/authentication/internal/service/settings"
	"github.com/nori-plugins/authentication/internal/service/user"
)

var ServiceSet = wire.NewSet(
	wire.Struct(new(auth.Params), "Config", "UserService", "AuthenticationLogService", "SessionService", "Transactor"),
	auth.New,
	wire.Struct(new(mfa_recovery_code.Params), "MfaRecoveryCodeRepository", "MfaRecoveryCodeHelper", "Config"),
	mfa_recovery_code.New,
	wire.Struct(new(mfa_secret.Params), "MfaSecretRepository", "UserRepository", "Config"),
	mfa_secret.New,
	wire.Struct(new(settings.Params), "SessionRepository", "UserRepository"),
	settings.New,
	wire.Struct(new(user.Params), "UserRepository", "Transactor", "Config"),
	user.New,
	wire.Struct(new(authentication_log.Params), "AuthenticationLogRepository", "Transactor"),
	authentication_log.New,
	wire.Struct(new(session.Params), "SessionRepository", "Transactor"),
	session.New,
	wire.Struct(new(service.Params), "AuthenticationService", "MfaRecoveryCodeService", "MfaSecretService", "SettingsService"),
	service.New,
)
