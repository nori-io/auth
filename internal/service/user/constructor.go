package user

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	userRepository repository.UserRepository
	config         config.Config
	transactor     transactor.Transactor
}

type Params struct {
	UserRepository repository.UserRepository
	Transactor     transactor.Transactor
	Сonfig         config.Config
}

func New(params Params) service.UserService {
	return &UserService{
		userRepository: params.UserRepository,
		config:         params.Сonfig,
		transactor:     params.Transactor,
	}
}
