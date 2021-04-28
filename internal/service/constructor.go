package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/authentication"
)

type Service struct {
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaTotpService         service.MfaTotpService
	settingsService        service.SettingsService
	userService            service.UserService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaTotpService         service.MfaTotpService
	SettingsService        service.SettingsService
	UserService            service.UserService
}

func New(params Params) authentication.Authentication {
	return &Service{
		authenticationService:  params.AuthenticationService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaTotpService:         params.MfaTotpService,
		settingsService:        params.SettingsService,
		userService:            params.UserService,
	}
}
