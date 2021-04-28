package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_totp"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service "github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaTotpService struct {
	mfaTotpRepository repository.MfaTotpRepository
	userService       service.UserService
	config            config.Config
	mfaTotpHelper     mfa_totp.MfaTotpHelper
	sessionService    service.SessionService
}

type Params struct {
	MfaTotpRepository repository.MfaTotpRepository
	UserService       service.UserService
	Config            config.Config
	mfaTotpHelper     mfa_totp.MfaTotpHelper
	sessionService    service.SessionService
}

func New(params Params) service.MfaTotpService {
	return &MfaTotpService{
		mfaTotpRepository: params.MfaTotpRepository,
		userService:       params.UserService,
		config:            params.Config,
		mfaTotpHelper:     params.mfaTotpHelper,
		sessionService:    params.sessionService,
	}
}
