package main

import (
	"context"

	p "github.com/nori-io/common/v4/pkg/domain/plugin"
	"github.com/nori-plugins/authentication/pkg/authentication"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/service"

	em "github.com/nori-io/common/v4/pkg/domain/enum/meta"

	"github.com/nori-io/common/v4/pkg/domain/config"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	m "github.com/nori-io/common/v4/pkg/meta"

	noriGorm "github.com/nori-io/interfaces/database/orm/gorm"
	conf "github.com/nori-plugins/authentication/internal/config"
)

func New() p.Plugin {
	return &plugin{}
}

type plugin struct {
	instance service.AuthenticationService
	config   conf.Config
	logger   logger.FieldLogger
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
		Interface:    authentication.AuthenticationInterface,
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
	p.config = conf.Config{
		CookiesName:                   config.String("cookies.name", "name of cookies for keeping session id"),
		CookiesPath:                   config.String("cookies.path", "path of cookies"),
		CookiesDomain:                 config.String("cookies.domain", "domain of cookies"),
		CookiesExpires:                config.Int64("cookies.expires", ""),
		CookiesMaxAge:                 config.Int("cookies.maxage", ""),
		CookiesSecure:                 config.Bool("cookies.secure", ""),
		CookiesHttpOnly:               config.Bool("cookies.httponly", ""),
		CookiesSameSite:               config.Int("cookies.samesite", ""),
		EmailVerification:             config.Bool("email.verification", "verification of email"),
		EmailActivationCodeTTLSeconds: config.Int64("email.activationcodettl", "time to live of email activation code in seconds"),
		UrlPrefix:                     config.String("url.prefix", "url prefix for all handlers"),
		UrlLogoutRedirect:             config.String("url.logout.redirect", "url for logout redirect"),
		MfaRecoveryCodePattern:        config.String("mfa.recoverycode.pattern", "pattern for mfa recovery codes"),
		MfaRecoveryCodeSymbols:        config.String("mfa.recoverycode.symbols", "symbols that use when mfa recovery code generating"),
		MfaRecoveryCodeLength:         config.Int("mfa.recoverycode.maxlength", "max length of mfaRecoveryCode"),
		MfaRecoveryCodeCount:          config.Int("mfa.recoverycode.count", "count of mfa recovery codes"),
		Issuer:                        config.String("mfa.issuer", "issuer"),
		PasswordBcryptCost:            config.Int("password.bcrypt.cost", "cost passed into GenerateFromPassword func"),
	}

	p.logger = log
	return nil
}

func (p plugin) Start(ctx context.Context, registry registry.Registry) error {
	config := p.config
	_, err := Initialize(registry, config, p.logger)
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
	//@todo actual sql code
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
