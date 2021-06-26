package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/service/user_log"
	"github.com/nori-plugins/authentication/pkg/authentication"
)

type Service struct {
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaTotpService         service.MfaTotpService
	resetPasswordService   service.ResetPasswordService
	sessionService         service.SessionService
	settingsService        service.SettingsService
	socialProviderService  service.SocialProvider
	userService            service.UserService
	userLogService         user_log.UserLogService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaTotpService         service.MfaTotpService
	ResetPasswordService   service.ResetPasswordService
	SessionService         service.SessionService
	SettingsService        service.SettingsService
	SocialProviderService  service.SocialProvider
	UserService            service.UserService
	UserLogService         user_log.UserLogService
}

func New(params Params) authentication.Authentication {
	return &Service{
		authenticationService:  params.AuthenticationService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaTotpService:         params.MfaTotpService,
		resetPasswordService:   params.ResetPasswordService,
		sessionService:         params.SessionService,
		settingsService:        params.SettingsService,
		socialProviderService:  params.SocialProviderService,
		userService:            params.UserService,
		userLogService:         params.UserLogService,
	}
}
