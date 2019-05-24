package main

import (
	"context"

	cfg "github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/meta"
	noriPlugin "github.com/nori-io/nori-common/plugin"

	"github.com/nori-io/authentication/service"
	"github.com/nori-io/authentication/service/database"
)

type plugin struct {
	instance service.Service
	config   *service.Config
}

var (
	Plugin plugin
)

func (p *plugin) Init(_ context.Context, configManager cfg.Manager) error {

	configManager.Register(p.Meta())
	cm := configManager.Register(p.Meta())
	p.config = &service.Config{
		Sub:                            cm.String("jwt.sub", "jwt.sub value"),
		Iss:                            cm.String("jwt.iss", "jwt.iss value"),
		UserType:                       cm.Slice("user.type", ",", "no"),
		UserTypeDefault:                cm.String("user.type_default", "user.type_default value"),
		UserRegistrationByPhoneNumber:  cm.Bool("user.registration_phone_number", "user.registration_phone_number value"),
		UserRegistrationByEmailAddress: cm.Bool("user.registration_email_address", "user.registration_email_address value"),
		UserMfaType:                    cm.String("user.mfa_type", "user.mfa_type value"),
		MailActivationTimeMinutes:      cm.UInt("mail.activation_time_minutes", "mail.activation_time_minutes value"),
		MailActivationCodeUsing:cm.Bool("mail.actication_code_using", "mail.actication_code_using"),
	}
	return nil
}

func (p *plugin) Start(_ context.Context, registry noriPlugin.Registry) error {

	if p.instance == nil {

		auth, err := registry.Auth()
		if err != nil {
			return err
		}

		cache, err := registry.Cache()
		if err != nil {
			return err
		}

		db, err := registry.Sql()
		if err != nil {
			return err
		}

		http, err := registry.Http()
		if err != nil {
			return err
		}

		transport, err := registry.HTTPTransport()
		if err != nil {
			return err
		}

		mail, err := registry.Mail()
		if err != nil {
			return err
		}

		session, err := registry.Session()
		if err != nil {
			return err
		}

		p.instance = service.NewService(
			auth,
			cache,
			p.config,
			database.DB(db.GetDB(), registry.Logger(p.Meta())),
			registry.Logger(p.Meta()),
			mail,
			session,
		)
		pluginParameters := service.PluginParameters{
			UserTypeParameter:              p.config.UserType(),
			UserTypeDefaultParameter:       p.config.UserTypeDefault(),
			UserRegistrationByPhoneNumber:  p.config.UserRegistrationByPhoneNumber(),
			UserRegistrationByEmailAddress: p.config.UserRegistrationByPhoneNumber(),
			UserMfaTypeParameter:           p.config.UserMfaType()}

		service.Transport(auth, transport, session,
			http, p.instance, registry.Logger(p.Meta()), pluginParameters)

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
			ID:      "nori/authentication",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{
			meta.Auth.Dependency("1.0.0"),
			meta.HTTP.Dependency("1.0.0"),
			meta.SQL.Dependency("1.0.0"),
			meta.Mail.Dependency("1.0.0"),
			meta.HTTPTransport.Dependency("1.0.0"),
		},
		Description: meta.Description{
			Name: "Nori: Authentication Interface",
		},
		Interface: meta.Custom,
		License: meta.License{
			Title: "",
			Type:  "LGPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		Tags: []string{"api"},
	}
}

func (p plugin) Install(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := registry.Sql()
	if err != nil {
		return err
	}
	db := database.DB(sql.GetDB(), registry.Logger(p.Meta()))
	db.CreateTables()
	return err
}

func (p plugin) UnInstall(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := registry.Sql()
	if err != nil {
		return err
	}
	db := database.DB(sql.GetDB(), registry.Logger(p.Meta()))
	db.DropTables()
	return err

}
