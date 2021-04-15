package settings

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SettingsService struct {
	sessionRepository repository.SessionRepository
	userService       service.UserService
	config            config.Config
	securityHelper    security.SecurityHelper
}

type Params struct {
	SessionRepository repository.SessionRepository
	UserService       service.UserService
	Config            config.Config
	SecurityHelper    security.SecurityHelper
}

func New(params Params) service.SettingsService {
	return &SettingsService{
		sessionRepository: params.SessionRepository,
		userService:       params.UserService,
		config:            params.Config,
		securityHelper:    params.SecurityHelper,
	}
}
