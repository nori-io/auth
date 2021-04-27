package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/totp"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service "github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaTotpService struct {
	mfaSecretRepository repository.MfaSecretRepository
	userService         service.UserService
	config              config.Config
	totpHelper          totp.TotpHelper
	sessionService      service.SessionService
}

type Params struct {
	MfaSecretRepository repository.MfaSecretRepository
	UserService         service.UserService
	Config              config.Config
	TotpHelper          totp.TotpHelper
	sessionService      service.SessionService
}

func New(params Params) service.MfaSecretService {
	return &MfaTotpService{
		mfaSecretRepository: params.MfaSecretRepository,
		userService:         params.UserService,
		config:              params.Config,
		totpHelper:          params.TotpHelper,
		sessionService:      params.sessionService,
	}
}
