package user

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	userRepository repository.UserRepository
	transactor     transactor.Transactor
	config         config.Config
}

type Params struct {
	UserRepository repository.UserRepository
	Transactor     transactor.Transactor
	Config         config.Config
}

func New(params Params) service.UserService {
	return &UserService{
		userRepository: params.UserRepository,
		transactor:     params.Transactor,
		config:         params.Config,
	}
}
