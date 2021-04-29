package user

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	authenticationLogService service.AuthenticationLogService
	userRepository           repository.UserRepository
	transactor               transactor.Transactor
	config                   config.Config
	securityHelper           security.SecurityHelper
}

type Params struct {
	authenticationLogService service.AuthenticationLogService
	UserRepository           repository.UserRepository
	Transactor               transactor.Transactor
	Config                   config.Config
	SecurityHelper           security.SecurityHelper
}

func New(params Params) service.UserService {
	return &UserService{
		authenticationLogService: params.authenticationLogService,
		userRepository:           params.UserRepository,
		transactor:               params.Transactor,
		config:                   params.Config,
		securityHelper:           params.SecurityHelper,
	}
}
