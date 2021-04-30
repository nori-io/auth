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
	userService              service.UserService
	sessionRepository        repository.SessionRepository
	securityHelper           security.SecurityHelper
	config                   config.Config
	transactor               transactor.Transactor
}

type Params struct {
	AuthenticationLogService service.AuthenticationLogService
	UserService              service.UserService
	SessionRepository        repository.SessionRepository
	SecurityHelper           security.SecurityHelper
	Config                   config.Config
	transactor               transactor.Transactor
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		authenticationLogService: params.AuthenticationLogService,
		userService:              params.UserService,
		sessionRepository:        params.SessionRepository,
		securityHelper:           params.SecurityHelper,
		config:                   params.Config,
		transactor:               params.transactor,
	}
}
