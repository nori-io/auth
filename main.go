package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nori-io/auth/service"
	"github.com/nori-io/auth/service/database"
	"github.com/nori-io/auth/service/database/sqlScripts"
	cfg "github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/meta"
	noriPlugin "github.com/nori-io/nori-common/plugin"
	"log"
)

type plugin struct {
	instance service.Service
}

var (
	Plugin plugin
	ctx = context.Background()
)

func (p *plugin) Init(_ context.Context, configManager cfg.Manager) error {
	configManager.Register(p.Meta())
	return nil
}

func (p *plugin) Start(_ context.Context, registry noriPlugin.Registry) error {
	if p.instance == nil {
		http, err := registry.Http()
		if err != nil {
			return err
		}

		db, err := registry.Sql()
		if err != nil {
			return err
		}


		p.instance = service.NewService(
			registry.Logger(p.Meta()),
			database.DB(db.GetDB()),
		)

		sql1, err := registry.Sql()
		if err != nil {
			return err
		}
		db1:= sql1.GetDB()

		tx, err := db1.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		_, execErr := tx.Exec(
			sqlScripts.CreateTableUsers)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)

		}
		_, execErr = tx.Exec(
			sqlScripts.CreateTableAuth)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}
		_, execErr = tx.Exec(
			sqlScripts.CreateTableAuthProviders)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}

		_, execErr = tx.Exec(
			sqlScripts.CreateTableAuthenticationHistory)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}

		_, execErr = tx.Exec(
			sqlScripts.CreateTableUserMfaPhone)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}

		_, execErr = tx.Exec(
			sqlScripts.CreateTableUsersMfaCode)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}
		_, execErr = tx.Exec(
			sqlScripts.CreateTableUserMfaSecret)
		if execErr != nil {
			_ = tx.Rollback()
			fmt.Println(execErr)
			log.Fatal(execErr)
		}


		if err := tx.Commit(); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}




		service.Transport(http, p.instance, registry.Logger(p.Meta()))




	}
	return nil
}

func (p *plugin) Stop(_ context.Context, _ noriPlugin.Registry) error {
	p.instance = nil
	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p plugin) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/authorization",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://noricms.com",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{
			meta.SQL.Dependency("1.0.0"),
			meta.HTTP.Dependency("1.0.0"),
		},
		Description: meta.Description{
			Name:        "NoriCMS Naive Posts Plugin",
			Description: "Naive Posts Plugin",
		},
		Interface: meta.Custom,
		License: meta.License{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		Tags: []string{"cms", "posts", "api"},
	}
}

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
