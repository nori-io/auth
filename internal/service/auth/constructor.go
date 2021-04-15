package auth

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AuthenticationService struct {
	config                   config.Config
	userService              service.UserService
	authenticationLogService service.AuthenticationLogService
	sessionService           service.SessionService
	mfaRecoveryCodeService   service.MfaRecoveryCodeService
	transactor               transactor.Transactor
	securityHelper           security.SecurityHelper
}

type Params struct {
	Config                   config.Config
	UserService              service.UserService
	AuthenticationLogService service.AuthenticationLogService
	SessionService           service.SessionService
	MfaRecoveryCodeService   service.MfaRecoveryCodeService
	Transactor               transactor.Transactor
	SecurityHelper           security.SecurityHelper
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		config:                   params.Config,
		userService:              params.UserService,
		authenticationLogService: params.AuthenticationLogService,
		sessionService:           params.SessionService,
		mfaRecoveryCodeService:   params.MfaRecoveryCodeService,
		transactor:               params.Transactor,
		securityHelper:           params.SecurityHelper,
	}
}
