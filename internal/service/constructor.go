package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/service/user_log"
	"github.com/nori-plugins/authentication/pkg/authentication"
)

type Service struct {
	authenticationService  service.AuthenticationService
	userLogService         user_log.UserLogService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaTotpService         service.MfaTotpService
	resetPasswordService   service.ResetPasswordService
	sessionService         service.SessionService
	settingsService        service.SettingsService
	socialProviderService  service.SocialProvider
	userService            service.UserService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	UserLogService         user_log.UserLogService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaTotpService         service.MfaTotpService
	ResetPasswordService   service.ResetPasswordService
	SessionService         service.SessionService
	SettingsService        service.SettingsService
	SocialProviderService  service.SocialProvider
	UserService            service.UserService
}

func New(params Params) authentication.Authentication {
	return &Service{
		authenticationService:  params.AuthenticationService,
		userLogService:         params.UserLogService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaTotpService:         params.MfaTotpService,
		resetPasswordService:   params.ResetPasswordService,
		sessionService:         params.SessionService,
		settingsService:        params.SettingsService,
		socialProviderService:  params.SocialProviderService,
		userService:            params.UserService,
	}
}
