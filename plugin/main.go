package main

import (
	"context"

	plugin2 "github.com/nori-io/common/v3/pkg/domain/plugin"

	"github.com/google/wire"

	"github.com/nori-io/authentication/internal/handler/http/authentication"

	"go.uber.org/dig"

	"github.com/jinzhu/gorm"

	noriHttp "github.com/nori-io/interfaces/nori/http"

	"github.com/nori-io/authentication/internal/domain/service"

	"github.com/nori-io/authentication/internal/repository/user"

	"github.com/nori-io/authentication/internal/service/auth"

	"github.com/nori-io/authentication/pkg"

	em "github.com/nori-io/common/v3/pkg/domain/enum/meta"

	"github.com/nori-io/common/v3/pkg/domain/config"
	"github.com/nori-io/common/v3/pkg/domain/logger"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	m "github.com/nori-io/common/v3/pkg/meta"
	s "github.com/nori-io/interfaces/nori/session"

	noriGorm "github.com/nori-io/interfaces/public/sql/gorm"
)

var Plugin plugin2.Plugin = pluginStruct{}

type pluginStruct struct {
	instance service.AuthenticationService
	config   conf
}

type conf struct {
	urlPrefix config.String
}

func (p pluginStruct) Meta() meta.Meta {
	return m.Meta{
		ID: m.ID{
			ID:      "",
			Version: "",
		},
		Author: m.Author{
			Name: "",
			URL:  "",
		},
		Dependencies: []meta.Dependency{},
		Description:  nil,
		Interface:    pkg.AuthenticationInterface,
		License:      nil,
		Links:        nil,
		Repository: m.Repository{
			Type: em.Git,
			URL:  "github.com/nori-io/http",
		},
		Tags: nil,
	}
}

func (p pluginStruct) Instance() interface{} {
	return p.instance
}

func (p pluginStruct) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = conf{
		urlPrefix: config.String("urlPrefix", "url prefix for all handlers"),
	}

	return nil
}

func (p pluginStruct) Start(ctx context.Context, registry registry.Registry) error {
	container := dig.New()
	container.Provide(registry)
	container.Provide(noriHttp.GetHttp)
	container.Provide(noriGorm.GetGorm)
	container.Provide(s.GetSession)
	container.Provide(user.New)
	container.Provide(authentication.New)
	container.Provide(auth.New)
	err := container.Invoke(func(router noriHttp.Http, handler *authentication.AuthHandler) {
		router.Get(p.config.urlPrefix()+"/signup", handler.SignUp)
		router.Get(p.config.urlPrefix()+"/signin", handler.SigIn)
		router.Get(p.config.urlPrefix()+"/signout", handler.SignOut)
	})
	if err != nil {
		return err
	}
	wire.NewSet(
		user.New,
		auth.New,
		authentication.New,
		interface{}(noriGorm.GetGorm(registry)),
		interface{}(s.GetSession(registry)),
		interface{}(noriHttp.GetHttp(registry)))

	/*db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}*/

	/*s, err := s.GetSession(registry)
	if err != nil {
		return err
	}*/

	// userRepo := user.New(db)

	// p.instance = auth.New(s, userRepo)

	/*httpServer, err := noriHttp.GetHttp(registry)
	if err != nil {
		return err
	}*/

	/*h := http.Handler{
		R:         httpServer,
		Auth:      p.instance,
		UrlPrefix: p.config.urlPrefix(),
	}*/

	// http.New(h)

	return nil
}

func (p pluginStruct) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p pluginStruct) Install(_ context.Context, registry registry.Registry) error {
	db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(`CREATE TABLE users(
		id bigserial PRIMARY KEY,
		email  VARCHAR (32) NOT NULL,
		password VARCHAR (32) NOT NULL,
		status   SMALLINT NOT NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
		);
		`).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (p pluginStruct) UnInstall(_ context.Context, registry registry.Registry) error {
	db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(`drop table users;
		`).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
