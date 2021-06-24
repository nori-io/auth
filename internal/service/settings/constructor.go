package settings

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SettingsService struct {
	userService       service.UserService
	userLogService    service.UserLogService
	sessionRepository repository.SessionRepository
	securityHelper    helper.SecurityHelper
	config            config.Config
	transactor        transactor.Transactor
}

type Params struct {
	UserService       service.UserService
	UserLogService    service.UserLogService
	SessionRepository repository.SessionRepository
	SecurityHelper    helper.SecurityHelper
	Config            config.Config
	Transactor        transactor.Transactor
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		userService:       params.UserService,
		userLogService:    params.UserLogService,
		sessionRepository: params.SessionRepository,
		securityHelper:    params.SecurityHelper,
		config:            params.Config,
		transactor:        params.Transactor,
	}
}
