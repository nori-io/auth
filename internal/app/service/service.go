package service

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/service"
	"github.com/nori-plugins/authentication/internal/service/auth"
	"github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/service/secret"
)

var ServiceSet = wire.NewSet(
	wire.Struct(new(mfa_recovery_code.ServiceParams), "MfaRecoveryCodeRepository", "MfaRecoveryCodeHelper", "Config"),
	auth.New,
	mfa_recovery_code.New,
	secret.New,
	service.New,
)
