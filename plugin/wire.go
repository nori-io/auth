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
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/user"
)

var set1 = wire.NewSet(
	noriGorm.GetGorm,
	session.GetSession,
	user.New,
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationService", "MfaRecoveryCodeService", "UrlPrefix"),
	noriHttp.GetHttp,
	authentication.New,
	mfa_recovery_code.New,
)

func Initialize(registry registry.Registry, urlPrefix string) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
