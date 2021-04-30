package auth

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AuthenticationService struct {
	authenticationLogService service.AuthenticationLogService
	mfaRecoveryCodeService   service.MfaRecoveryCodeService
	mfaTotpService           service.MfaTotpService
	sessionService           service.SessionService
	socialProviderService    service.SocialProvider
	userService              service.UserService
	securityHelper           security.SecurityHelper
	config                   config.Config
	transactor               transactor.Transactor
}

type Params struct {
	AuthenticationLogService service.AuthenticationLogService
	MfaRecoveryCodeService   service.MfaRecoveryCodeService
	MfaTotpService           service.MfaTotpService
	SessionService           service.SessionService
	SocialProviderService    service.SocialProvider
	UserService              service.UserService
	SecurityHelper           security.SecurityHelper
	Config                   config.Config
	Transactor               transactor.Transactor
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		authenticationLogService: params.AuthenticationLogService,
		mfaRecoveryCodeService:   params.MfaRecoveryCodeService,
		mfaTotpService:           params.MfaTotpService,
		sessionService:           params.SessionService,
		socialProviderService:    params.SocialProviderService,
		userService:              params.UserService,
		securityHelper:           params.SecurityHelper,
		config:                   params.Config,
		transactor:               params.Transactor,
	}
}
