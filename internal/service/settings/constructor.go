package settings

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SettingsService struct {
	authenticationLogService service.AuthenticationLogService
	sessionRepository        repository.SessionRepository
	userService              service.UserService
	config                   config.Config
	securityHelper           security.SecurityHelper
	transactor               transactor.Transactor
}

type Params struct {
	authenticationLogService service.AuthenticationLogService
	SessionRepository        repository.SessionRepository
	UserService              service.UserService
	Config                   config.Config
	SecurityHelper           security.SecurityHelper
	Transactor               transactor.Transactor
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		authenticationLogService: params.authenticationLogService,
		sessionRepository:        params.SessionRepository,
		userService:              params.UserService,
		config:                   params.Config,
		securityHelper:           params.SecurityHelper,
		transactor:               params.Transactor,
	}
}
