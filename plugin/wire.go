///+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/config"
	httpHandler "github.com/nori-plugins/authentication/internal/handler/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_codes"
	mfaRecoveryCodeRepository "github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/user"
	servAuth "github.com/nori-plugins/authentication/internal/service/auth"
	servMfaRecoveryCode "github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
)

var set1 = wire.NewSet(

	wire.Struct(new(servMfaRecoveryCode.ServiceParams), "MfaRecoveryCodeRepository", "MfaRecoveryCodeHelper", "Config"),
	noriGorm.GetGorm,
	mfaRecoveryCodeRepository.New,
	mfaRecoveryCodeHelper.New,
	session.GetSession,
	user.New,
	servAuth.New,
	servMfaRecoveryCode.New,
	authentication.New,
	mfa_recovery_code.New,
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationService",
		"MfaRecoveryCodeService", "UrlPrefix", "AuthenticationHandler", "MfaRecoveryCodeHandler"),
	noriHttp.GetHttp)

func Initialize(registry registry.Registry, config config.Config) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
