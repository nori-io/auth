///+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/app"
	"github.com/nori-plugins/authentication/internal/config"
	httpHandler "github.com/nori-plugins/authentication/internal/handler/http"
)

var set1 = wire.NewSet(
	noriGorm.GetGorm,
	session.GetSession,
	noriHttp.GetHttp)

func Initialize(registry registry.Registry, config config.Config) (*httpHandler.Handler, error) {
	wire.Build(app.AppSet, set1)
	return &httpHandler.Handler{}, nil
}
