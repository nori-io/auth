package service

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/service"
	"github.com/nori-plugins/authentication/internal/service/auth"
	"github.com/nori-plugins/authentication/internal/service/authentication_log"
	"github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/service/mfa_totp"
	"github.com/nori-plugins/authentication/internal/service/session"
	"github.com/nori-plugins/authentication/internal/service/settings"
	"github.com/nori-plugins/authentication/internal/service/social_provider"
	"github.com/nori-plugins/authentication/internal/service/user"
)

var ServiceSet = wire.NewSet(
	wire.Struct(new(auth.Params), "Config", "UserService", "AuthenticationLogService", "SessionService", "Transactor", "SecurityHelper"),
	auth.New,
	wire.Struct(new(mfa_recovery_code.Params), "MfaRecoveryCodeRepository", "MfaRecoveryCodeHelper", "Config"),
	mfa_recovery_code.New,
	wire.Struct(new(mfa_totp.Params), "MfaTotpRepository", "UserService", "Config"),
	mfa_totp.New,
	wire.Struct(new(settings.Params), "SessionRepository", "UserService", "SecurityHelper"),
	settings.New,
	wire.Struct(new(user.Params), "UserRepository", "Transactor", "Config", "SecurityHelper"),
	user.New,
	wire.Struct(new(authentication_log.Params), "AuthenticationLogRepository", "Transactor"),
	authentication_log.New,
	wire.Struct(new(session.Params), "SessionRepository", "Transactor"),
	session.New,
	wire.Struct(new(social_provider.Params), "SocialProviderRepository"),
	social_provider.New,
	wire.Struct(new(service.Params), "AuthenticationService", "MfaRecoveryCodeService", "MfaTotpService", "SettingsService"),
	service.New,
)
