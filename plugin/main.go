package main

import (
	"context"

	authentication2 "github.com/nori-plugins/authentication/pkg/authentication"

	p "github.com/nori-io/common/v4/pkg/domain/plugin"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/service"

	em "github.com/nori-io/common/v4/pkg/domain/enum/meta"

	"github.com/nori-io/common/v4/pkg/domain/config"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	m "github.com/nori-io/common/v4/pkg/meta"

	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	instance service.AuthenticationService
	config   Conf
}

type Conf struct {
	urlPrefix                config.String
	MfaRecoveryCodePattern   config.String
	MfaRecoveryCodeSymbols   config.String
	MfaRecoveryCodeMaxLength config.Int
	MfaRecoveryCodeCount     config.Int
	Issuer                   config.String
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
		Interface:    authentication2.AuthenticationInterface,
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
	p.config = Conf{
		urlPrefix:                config.String("urlprefix", "url prefix for all handlers"),
		MfaRecoveryCodePattern:   config.String("mfa.recoverycode.pattern", "pattern for mfa recovery codes"),
		MfaRecoveryCodeSymbols:   config.String("mfa.recoverycode.symbols", "symbols that use when mfa recovery code generating"),
		MfaRecoveryCodeMaxLength: config.Int("mfa.recoverycode.maxlength", "max length of mfaRecoveryCode"),
		MfaRecoveryCodeCount:     config.Int("mfa.recoverycode.count", "count of mfa recovery codes"),
		Issuer:                   config.String("mfa.issuer", "issuer"),
	}

	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {
	_, err := Initialize(registry, p.config.urlPrefix(), p.config.MfaRecoveryCodeCount())
	return err
}

func (p plugin) Stop(ctx context.Context, registry registry.Registry) error {
	return nil
}

func (p plugin) Install(_ context.Context, registry registry.Registry) error {
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

func (p plugin) UnInstall(_ context.Context, registry registry.Registry) error {
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
