package main

import (
	"context"

	"github.com/nori-io/authentication/internal/domain/service"

	"github.com/nori-io/authentication/internal/repository/user"

	"github.com/nori-io/authentication/internal/service/auth"

	"github.com/jinzhu/gorm"

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
	db       *gorm.DB
	instance service.AuthenticationService
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
	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {

	db, _ := noriGorm.GetGorm(registry)
	s, _ := s.GetSession(registry)
	userRepo := user.New(db)
	p.instance = auth.New(s, userRepo)

	/*if p.instance == nil {

	http, err := registry.Http()
	if err != nil {
		return err
	}

	session, err := registry.Session()
	if err != nil {
		return err
	}

	db, err := registry.Sql()
	if err != nil {
		return err
	}*/

	/*p.instance = service.NewService(
		auth,
		session,
		p.config,
		registry.Logger(p.Meta()),
		database.DB(db.GetDB()),
	)
	service.Transport(auth, transport, session,
		http, p.instance, registry.Logger(p.Meta()))*/

	/*		sql1, err := registry.Sql()
			if err != nil {
				return err
			}
			db1 := sql1.GetDB()

			tx, err := db1.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

			if err != nil {
				log.Fatal(err)
			}
			_, execErr := tx.Exec(
				sqlScripts.CreateTableUsers)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)

			}
			_, execErr = tx.Exec(
				sqlScripts.CreateTableAuth)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}
			_, execErr = tx.Exec(
				sqlScripts.CreateTableAuthProviders)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}

			_, execErr = tx.Exec(
				sqlScripts.CreateTableAuthenticationHistory)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}

			_, execErr = tx.Exec(
				sqlScripts.CreateTableUserMfaPhone)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}

			_, execErr = tx.Exec(
				sqlScripts.CreateTableUsersMfaCode)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}
			_, execErr = tx.Exec(
				sqlScripts.CreateTableUserMfaSecret)
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}

			if err := tx.Commit(); err != nil {
				log.Fatal(err)
			}

			/*service.Transport(auth, transport, session,
				http, p.instance, registry.Logger(p.Meta()))*/
	//}
	return nil
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

/*
func (p plugin) Install(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := registry.Sql()
	if err != nil {
		return err
	}
	db := sql.GetDB()
	_, err = db.Exec(sqlScripts.CreateTableUsersMfaCode)
	return err
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
*/
