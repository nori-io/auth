package auth

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AuthenticationService struct {
	userLogService         service.UserLogService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	mfaTotpService         service.MfaTotpService
	sessionService         service.SessionService
	socialProviderService  service.SocialProvider
	userService            service.UserService
	securityHelper         security.SecurityHelper
	config                 config.Config
	transactor             transactor.Transactor
}

type Params struct {
	UserLogService         service.UserLogService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
	MfaTotpService         service.MfaTotpService
	SessionService         service.SessionService
	SocialProviderService  service.SocialProvider
	UserService            service.UserService
	SecurityHelper         security.SecurityHelper
	Config                 config.Config
	Transactor             transactor.Transactor
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		userLogService:         params.UserLogService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
		mfaTotpService:         params.MfaTotpService,
		sessionService:         params.SessionService,
		socialProviderService:  params.SocialProviderService,
		userService:            params.UserService,
		securityHelper:         params.SecurityHelper,
		config:                 params.Config,
		transactor:             params.Transactor,
	}
}
