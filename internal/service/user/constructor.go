package user

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserService struct {
	userRepository repository.UserRepository
	db             *gorm.DB
	config         config.Config
	transactor     transactor.Transactor
}

type Params struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	config         config.Config
}

func New(params Params) service.UserService {
	return &UserService{
		userRepository: params.UserRepository,
		db:             params.DB,
		config:         params.config,
	}
}
