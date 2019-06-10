package main

import (
	"context"

	cfg "github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/meta"
	noriPlugin "github.com/nori-io/nori-common/plugin"
	"github.com/nori-io/nori-interfaces/interfaces"

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
		Sub:                                cm.String("jwt.sub", "jwt.sub value"),
		Iss:                                cm.String("jwt.iss", "jwt.iss value"),
		UserType:                           cm.Slice("user.type", ",", "no"),
		UserTypeDefault:                    cm.String("user.type_default", "user.type_default value"),
		UserRegistrationByPhoneNumber:      cm.Bool("user.registration_phone_number", "user.registration_phone_number value"),
		UserRegistrationByEmailAddress:     cm.Bool("user.registration_email_address", "user.registration_email_address value"),
		UserMfaType:                        cm.String("user.mfa_type", "user.mfa_type value"),
		ActivationTimeForActivationMinutes: cm.UInt("activation.time_for_activation_minutes", "activation.time_for_activation_minutes value"),
		ActivationCode:                     cm.Bool("activation.code", "activation.code value"),
		Oath2ProvidersVKClientKey:          cm.String("oath2.providers.vk.client_key", "oath2.providers.vk.client_key value"),
		Oath2ProvidersVKClientSecret:       cm.String("oath2.providers.vk.client_secret", "oath2.providers.vk.client_secret value"),
		Oath2ProvidersVKRedirectUrl:        cm.String("oath2.providers.vk.redirect_url", "oath2.providers.vk.client_redirect_url"),
		Oath2SessionSecret:                 cm.String("oath2.session_secret", "oath2.session_secret"),
	}
	return nil
}

func (p *plugin) Start(_ context.Context, registry noriPlugin.Registry) error {

	if p.instance == nil {

		auth, err := interfaces.GetAuth(registry)
		if err != nil {
			return err
		}
		cache, err := interfaces.GetCache(registry)
		if err != nil {
			return err
		}

		db, err := interfaces.GetSQL(registry)
		if err != nil {
			return err
		}

		http, err := interfaces.GetHttp(registry)
		if err != nil {
			return err
		}

		transport, err := interfaces.GetHttpTransport(registry)
		if err != nil {
			return err
		}

		mail, err := interfaces.GetMail(registry)
		if err != nil {
			return err
		}

		session, err := interfaces.GetSession(registry)
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
			UserTypeParameter:                  p.config.UserType(),
			UserTypeDefaultParameter:           p.config.UserTypeDefault(),
			UserRegistrationByPhoneNumber:      p.config.UserRegistrationByPhoneNumber(),
			UserRegistrationByEmailAddress:     p.config.UserRegistrationByPhoneNumber(),
			UserMfaTypeParameter:               p.config.UserMfaType(),
			ActivationCode:                     p.config.ActivationCode(),
			ActivationTimeForActivationMinutes: p.config.ActivationTimeForActivationMinutes(),
			Oath2ProvidersVKClientKey:          p.config.Oath2ProvidersVKClientKey(),
			Oath2ProvidersVKClientSecret:       p.config.Oath2ProvidersVKClientSecret(),
			Oath2ProvidersVKRedirectUrl:        p.config.Oath2ProvidersVKRedirectUrl(),
			Oath2SessionSecret:                 p.config.Oath2SessionSecret(),
		}

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
			interfaces.AuthInterface.Dependency("1.0.0"),
			interfaces.MailInterface.Dependency("1.0.0"),
			interfaces.CacheInterface.Dependency("1.0.0"),
			interfaces.HttpInterface.Dependency("1.0.0"),
			interfaces.HttpTransportInterface.Dependency("1.0.0"),
			interfaces.SQLInterface.Dependency("1.0.0"),
			interfaces.SessionInterface.Dependency("1.0.0"),
		},
		Description: meta.Description{
			Name: "Nori: Authentication Interface",
		},
		Interface: meta.Interface("Custom"),
		License: meta.License{
			Title: "",
			Type:  "LGPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		Tags: []string{"api"},
	}
}

func (p plugin) Install(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := interfaces.GetSQL(registry)
	if err != nil {
		return err
	}
	db := database.DB(sql.GetDB(), registry.Logger(p.Meta()))
	db.CreateTables()
	return err
}

func (p plugin) UnInstall(_ context.Context, registry noriPlugin.Registry) error {
	sql, err := interfaces.GetSQL(registry)
	if err != nil {
		return err
	}
	db := database.DB(sql.GetDB(), registry.Logger(p.Meta()))
	db.DropTables()
	return err

}
