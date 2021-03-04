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
)

var set1 = wire.NewSet(
	noriGorm.GetGorm,
	session.GetSession,
	wire.Struct(new(httpHandler.Handler), "R", "AuthenticationService",
		"MfaRecoveryCodeService", "UrlPrefix", "AuthenticationHandler", "MfaRecoveryCodeHandler"),
	noriHttp.GetHttp)

func Initialize(registry registry.Registry, config config.Config) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
