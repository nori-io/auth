package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/authentication"
)

type Service struct {
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaSecretService       service.MfaSecretService
	settingsService        service.SettingsService
	userService            service.UserService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaSecretService       service.MfaSecretService
	SettingsService        service.SettingsService
	UserService            service.UserService
}

func New(params Params) authentication.Authentication {
	return &Service{
		authenticationService:  params.AuthenticationService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaSecretService:       params.MfaSecretService,
		settingsService:        params.SettingsService,
		userService:            params.UserService,
	}
}
