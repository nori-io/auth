package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_totp"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service "github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type MfaTotpService struct {
	sessionService    service.SessionService
	userService       service.UserService
	userLogService    service.UserLogService
	mfaTotpRepository repository.MfaTotpRepository
	mfaTotpHelper     mfa_totp.MfaTotpHelper
	config            config.Config
	transactor        transactor.Transactor
}

type Params struct {
	SessionService    service.SessionService
	UserService       service.UserService
	UserLogService    service.UserLogService
	MfaTotpRepository repository.MfaTotpRepository
	MfaTotpHelper     mfa_totp.MfaTotpHelper
	Config            config.Config
	Transactor        transactor.Transactor
}

func New(params Params) service.MfaTotpService {
	return &MfaTotpService{
		sessionService:    params.SessionService,
		userService:       params.UserService,
		userLogService:    params.UserLogService,
		mfaTotpRepository: params.MfaTotpRepository,
		mfaTotpHelper:     params.MfaTotpHelper,
		config:            params.Config,
		transactor:        params.Transactor,
	}
}
