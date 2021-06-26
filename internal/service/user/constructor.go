package user

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	userLogService service.UserLogService
	userRepository repository.UserRepository
	securityHelper helper.SecurityHelper
	transactor     transactor.Transactor
	config         config.Config
}

type Params struct {
	UserLogService service.UserLogService
	UserRepository repository.UserRepository
	SecurityHelper helper.SecurityHelper
	Transactor     transactor.Transactor
	Config         config.Config
}

func New(params Params) service.UserService {
	return &UserService{
		userLogService: params.UserLogService,
		userRepository: params.UserRepository,
		securityHelper: params.SecurityHelper,
		transactor:     params.Transactor,
		config:         params.Config,
	}
}
