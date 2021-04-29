package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_totp"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service "github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type MfaTotpService struct {
	authenticationLogService service.AuthenticationLogService
	mfaTotpRepository        repository.MfaTotpRepository
	userService              service.UserService
	config                   config.Config
	mfaTotpHelper            mfa_totp.MfaTotpHelper
	sessionService           service.SessionService
	transactor               transactor.Transactor
}

type Params struct {
	authenticationLogService service.AuthenticationLogService
	MfaTotpRepository        repository.MfaTotpRepository
	UserService              service.UserService
	Config                   config.Config
	mfaTotpHelper            mfa_totp.MfaTotpHelper
	sessionService           service.SessionService
	transactor               transactor.Transactor
}

func New(params Params) service.MfaTotpService {
	return &MfaTotpService{
		authenticationLogService: params.authenticationLogService,
		mfaTotpRepository:        params.MfaTotpRepository,
		userService:              params.UserService,
		config:                   params.Config,
		mfaTotpHelper:            params.mfaTotpHelper,
		sessionService:           params.sessionService,
		transactor:               params.transactor,
	}
}
