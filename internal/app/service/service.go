package service

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/service"
	"github.com/nori-plugins/authentication/internal/service/auth"
	"github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/service/mfa_secret"
)

var ServiceSet = wire.NewSet(
	wire.Struct(new(auth.Params), "authenticationLogRepository", "db"),
	auth.New,
	wire.Struct(new(mfa_recovery_code.Params), "mfaRecoveryCodeRepository", "MfaRecoveryCodeHelper", "config"),
	mfa_recovery_code.New,
	wire.Struct(new(mfa_secret.Params), "MfaSecretRepository", "userRepository", "config"),
	mfa_secret.New,
	wire.Struct(new(service.Params), "AuthenticationService", "MfaRecoveryCodeService", "MfaSecretService"),
	service.New,
)
