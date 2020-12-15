package main

import (
	"github.com/google/wire"
	httpHandler "github.com/nori-io/authentication/internal/handler/http"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/authentication/internal/repository/user"
	"github.com/nori-io/authentication/internal/service/auth"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	"github.com/nori-io/interfaces/nori/session"
)

/*type registryManager struct {
	log           *logrus.Logger
	plugins       *PluginList
	interfaces    map[meta.Interface]meta.ID
	configManager config.Manager
	registry      plugin.Registry
}*/

// var registry1 registry.Registry

var set1 = wire.NewSet(
	auth.New,
	authentication.New,
	noriGorm.GetGorm,
	session.GetSession,
	http.GetHttp,
	user.New,
	httpHandler.New,
)

func Initialize(registry registry.Registry) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
