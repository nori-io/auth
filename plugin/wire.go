package main

import (
	"github.com/google/wire"
	httpHandler "github.com/nori-io/authentication/internal/handler/http"
	"github.com/nori-io/authentication/internal/repository/user"
	"github.com/nori-io/authentication/internal/service/auth"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	noriHttp "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-io/interfaces/nori/session"
)

var set1 = wire.NewSet(
	auth.New,
	// authentication.New,
	noriGorm.GetGorm,
	session.GetSession,
	user.New,
	wire.Struct(new(httpHandler.Handler), "R", "Auth", "UrlPrefix"),
	noriHttp.GetHttp,
)

func Initialize(registry registry.Registry, urlPrefix string) (*httpHandler.Handler, error) {
	wire.Build(set1)
	return &httpHandler.Handler{}, nil
}
