package main

import (
	"context"
	cfg "github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/meta"
	noriPlugin "github.com/nori-io/nori-common/plugin"
	"github.com/nori-io/auth/service"
	"github.com/nori-io/auth/service/database"
)

type plugin struct {
	instance service.Service
}

var (
	Plugin plugin
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
		service.Transport(			http, p.instance, registry.Logger(p.Meta()))
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
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS articles (
  		id INT AUTO_INCREMENT PRIMARY KEY,
  		title VARCHAR(255) NOT NULL,
  		body LONGTEXT NOT NULL,
  		state SET('draft', 'public', 'password', 'deleted') NOT NULL,
  		meta_description VARCHAR(255) NOT NULL,
 	    tags TEXT NOT NULL,
 		UNIQUE INDEX id_UNIQUE (id ASC) INVISIBLE,
  		INDEX state_idx (state ASC) INVISIBLE,
  		ENGINE = InnoDB;

		CREATE TABLE IF NOT EXISTS comments (
  		id INT AUTO_INCREMENT PRIMARY KEY,
  		parent_id INT NULL,
  		post_id INT NULL,
  		message TEXT NOT NULL,
  		created INT NOT NULL,
  		state SET('deleted', 'blocked', 'public') NOT NULL,
  		article_id_fk INT NOT NULL,
  		UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
  		INDEX state_idx (state ASC) INVISIBLE,
  		CONSTRAINT article_id_fk
		FOREIGN KEY (id)
    	REFERENCES db_articles.articles (id)
    	ON DELETE CASCADE
    	ON UPDATE CASCADE)
		ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS users (
		    id    BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name  varchar(255) NOT NULL,

    		email 			varchar(255) NOT NULL UNIQUE,
    		email_verified  tinyint NOT NULL DEFAULT FALSE,

    		phone              varchar(20),
    		phone_country_code varchar(5),
    		phone_verified     boolean NOT NULL DEFAULT FALSE,

    		salt               varchar(65) NOT NULL,
    		password           varchar(255) NOT NULL,

    		state              varchar(16) NOT NULL DEFAULT 'active',

    		mfa_enabled        tinyint NOT NULL DEFAULT 0,

    		created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    		updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)  ENGINE=InnoDB;
	`)
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
