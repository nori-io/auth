package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/service/authentication_log"
	"github.com/nori-plugins/authentication/pkg/authentication"
)

type Service struct {
	authenticationService    service.AuthenticationService
	authenticationLogService authentication_log.AuthenticationLogService
	mfaRecoveryCodeService   service.MfaRecoveryCodeService
	mfaTotpService           service.MfaTotpService
	resetPasswordService     service.ResetPasswordService
	sessionService           service.SessionService
	settingsService          service.SettingsService
	socialProviderService    service.SocialProvider
	userService              service.UserService
}

type Params struct {
	AuthenticationService    service.AuthenticationService
	AuthenticationLogService authentication_log.AuthenticationLogService
	MfaRecoveryCodeService   service.MfaRecoveryCodeService
	MfaTotpService           service.MfaTotpService
	ResetPasswordService     service.ResetPasswordService
	SessionService           service.SessionService
	SettingsService          service.SettingsService
	SocialProviderService    service.SocialProvider
	UserService              service.UserService
}

func New(params Params) authentication.Authentication {
	return &Service{
		authenticationService:    params.AuthenticationService,
		authenticationLogService: params.AuthenticationLogService,
		mfaRecoveryCodeService:   params.MfaRecoveryCodeService,
		mfaTotpService:           params.MfaTotpService,
		resetPasswordService:     params.ResetPasswordService,
		sessionService:           params.SessionService,
		settingsService:          params.SettingsService,
		socialProviderService:    params.SocialProviderService,
		userService:              params.UserService,
	}
}
