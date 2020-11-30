package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/common/v3/pkg/domain/logger"
	"github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/nori-io/common/v3/pkg/domain/registry"
	noriGorm "github.com/nori-io/interfaces/public/sql/gorm"
	m "github.com/nori-io/common/v3/pkg/meta"
	p "github.com/nori-io/common/v3/pkg/domain/plugin"
	"github.com/nori-io/common/v3/pkg/domain/config"
)
var (
	Plugin p.Plugin = plugin{}
)


type plugin struct {
	db *gorm.DB
	config conf

}

type conf struct {
	Sub config.String
	Iss config.String

}

func (p plugin) Meta() meta.Meta {
	return m.Meta{
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
			meta.HTTP.Dependency("1.0.0"),
			meta.SQL.Dependency("1.0.0"),
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

func (p plugin) Instance() interface{} {
	return p.db
}


func (p plugin) Init(ctx context.Context, config config.Config, log logger.FieldLogger) error {
	p.config = conf{
		Sub: config.String("jwt.sub", "jwt.sub value"),
		Iss: config.String("jwt.iss", "jwt.iss value"),
	}
	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {

	p.db, err := noriGorm.GetGorm(registry)

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