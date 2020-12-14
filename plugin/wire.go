package main

import (
	"github.com/google/wire"
	"github.com/nori-io/authentication/internal/domain/repository"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/authentication/internal/repository/user"
	"github.com/nori-io/authentication/internal/repository/user/postgres"
	"github.com/nori-io/authentication/internal/service/auth"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	s "github.com/nori-io/interfaces/nori/session"
	noriGorm "github.com/nori-io/interfaces/public/sql/gorm"
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
	s.GetSession,
	noriHttp.GetHttp,
	user.New,
)

func Initialize(registry registry.Registry) (repository.UserRepository, error) {
	wire.Build(set1)
	return &postgres.UserRepository{}, nil
}
