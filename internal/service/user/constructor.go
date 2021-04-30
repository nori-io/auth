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
	securityHelper           security.SecurityHelper
	transactor               transactor.Transactor
	config                   config.Config
}

type Params struct {
	authenticationLogService service.AuthenticationLogService
	UserRepository           repository.UserRepository
	SecurityHelper           security.SecurityHelper
	Transactor               transactor.Transactor
	Config                   config.Config
}

func New(params Params) service.UserService {
	return &UserService{
		authenticationLogService: params.authenticationLogService,
		userRepository:           params.UserRepository,
		securityHelper:           params.SecurityHelper,
		transactor:               params.Transactor,
		config:                   params.Config,
	}
}
