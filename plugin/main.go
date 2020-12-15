package main

import (
	"context"

	plugin2 "github.com/nori-io/common/v4/pkg/domain/plugin"

	//"go.uber.org/dig"

	"github.com/jinzhu/gorm"

	"github.com/nori-io/authentication/internal/domain/service"

	"github.com/nori-io/authentication/pkg"

	em "github.com/nori-io/common/v4/pkg/domain/enum/meta"

	// httpHandler "github.com/nori-io/authentication/internal/handler/http"
	"github.com/nori-io/common/v4/pkg/domain/config"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	m "github.com/nori-io/common/v4/pkg/meta"

	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
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
	/*httpServer, err := noriHttp.GetHttp(registry)
	if err != nil {
		return err
	}

	h := httpHandler.Handler{
		R:         httpServer,
		Auth:      p.instance,
		UrlPrefix: p.config.urlPrefix(),
	}*/

	Initialize(registry, p.config.urlPrefix())

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
