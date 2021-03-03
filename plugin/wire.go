///+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-io/interfaces/nori/session"
	httpHandler "github.com/nori-plugins/authentication/internal/handler/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"

	"github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/user"
	servAuth "github.com/nori-plugins/authentication/internal/service/auth"
	servMfaRecoveryCode "github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
)

var set1 = wire.NewSet(
	noriGorm.GetGorm,
	session.GetSession,
	user.New,
	servAuth.New,
	servMfaRecoveryCode.New,
	authentication.New,
	mfa_recovery_code.New,
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationService",
		"MfaRecoveryCodeService", "UrlPrefix", "AuthenticationHandler", "MfaRecoveryCodeHandler"),
	noriHttp.GetHttp,
)

func Initialize(registry registry.Registry, urlPrefix string) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
