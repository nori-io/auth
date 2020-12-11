package main

import (
	"github.com/google/wire"
	"github.com/nori-io/authentication/internal/domain/repository"
	"github.com/nori-io/authentication/internal/handler/http/authentication"
	"github.com/nori-io/authentication/internal/repository/user"
	"github.com/nori-io/authentication/internal/repository/user/postgres"
	"github.com/nori-io/authentication/internal/service/auth"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	s "github.com/nori-io/interfaces/nori/session"
	noriGorm "github.com/nori-io/interfaces/public/sql/gorm"
)

var set1 = wire.NewSet(
	user.New,
	auth.New,
	authentication.New,
	noriGorm.GetGorm,
	s.GetSession,
	noriHttp.GetHttp,
)

var registry2 registry.Registry

func Initialize() (repository.UserRepository, error) {
	wire.Build(set1)
	return &postgres.UserRepository{}, nil
}
