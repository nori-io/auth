package user

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	userRepository repository.UserRepository
	transactor     transactor.Transactor
	config         config.Config
	securityHelper security.SecurityHelper
}

type Params struct {
	UserRepository repository.UserRepository
	Transactor     transactor.Transactor
	Config         config.Config
	SecurityHelper security.SecurityHelper
}

func New(params Params) service.UserService {
	return &UserService{
		userRepository: params.UserRepository,
		transactor:     params.Transactor,
		config:         params.Config,
		securityHelper: params.SecurityHelper,
	}
}
