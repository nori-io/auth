package auth

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AuthenticationService struct {
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaTotpService         service.MfaTotpService
	sessionService         service.SessionService
	socialProviderService  service.SocialProvider
	userService            service.UserService
	userLogService         service.UserLogService
	securityHelper         security.SecurityHelper
	config                 config.Config
	transactor             transactor.Transactor
}

type Params struct {
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaTotpService         service.MfaTotpService
	SessionService         service.SessionService
	SocialProviderService  service.SocialProvider
	UserService            service.UserService
	UserLogService         service.UserLogService
	SecurityHelper         security.SecurityHelper
	Config                 config.Config
	Transactor             transactor.Transactor
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaTotpService:         params.MfaTotpService,
		sessionService:         params.SessionService,
		socialProviderService:  params.SocialProviderService,
		userService:            params.UserService,
		userLogService:         params.UserLogService,
		securityHelper:         params.SecurityHelper,
		config:                 params.Config,
		transactor:             params.Transactor,
	}
}
