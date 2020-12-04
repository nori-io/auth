package main

import (
	"context"

	"github.com/nori-io/authentication/internal/domain/service"

	"github.com/nori-io/authentication/internal/repository/user"

	"github.com/nori-io/authentication/internal/service/auth"

	"github.com/nori-io/authentication/pkg"
	em "github.com/nori-io/common/v3/pkg/domain/enum/meta"

	"github.com/nori-io/common/v3/pkg/domain/config"
	"github.com/nori-io/common/v3/pkg/domain/logger"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	p "github.com/nori-io/common/v3/pkg/domain/plugin"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	m "github.com/nori-io/common/v3/pkg/meta"
	s "github.com/nori-io/interfaces/nori/session"

	noriGorm "github.com/nori-io/interfaces/public/sql/gorm"
)

var (
	Plugin p.Plugin = plugin{}
)

type plugin struct {
	instance service.AuthenticationService
	config   conf
}

type conf struct {
	urlPrefix config.String
}

func (p plugin) Meta() meta.Meta {
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

func (p plugin) Instance() interface{} {
	return p.instance
}

func (p plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = conf{
		urlPrefix: config.String("urlPrefix", "url prefix for all handlers"),
	}

	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {

	db, _ := noriGorm.GetGorm(registry)
	s, _ := s.GetSession(registry)
	userRepo := user.New(db)
	p.instance = auth.New(s, userRepo)

	return nil
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Install(_ context.Context, registry registry.Registry) error {
	db, err := noriGorm.GetGorm(registry)
	if err != nil {
		return err
	}
	db.Exec("`CREATE TABLE users\n" +
		"(id bigserial PRIMARY KEY,\n" +
		" email  VARCHAR (32) NOT NULL,\n" +
		" password VARCHAR (32) NOT NULL,\n" +
		" status   SMALLINT NOT NULL,\n " +
		"created_at TIMESTAMP,\n " +
		"updated_at TIMESTAMP);`")
	return nil
}

func (p plugin) UnInstall(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := registry.Sql()
	if err != nil {
		return err
	}
	db := sql.GetDB()
	_, err = db.Exec(`
		drop table articles;
		drop table comments;
		`)
	return err
}
