package reset_password

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type ResetPasswordService struct {
	userService             service.UserService
	userLogService          service.UserLogService
	resetPasswordRepository repository.ResetPasswordRepository
	securityHelper          helper.SecurityHelper
	config                  config.Config
	transactor              transactor.Transactor
}

type Params struct {
	UserService             service.UserService
	UserLogService          service.UserLogService
	ResetPasswordRepository repository.ResetPasswordRepository
	SecurityHelper          helper.SecurityHelper
	Config                  config.Config
	Transactor              transactor.Transactor
}

func New(params Params) service.ResetPasswordService {
	return &ResetPasswordService{
		userService:             params.UserService,
		userLogService:          params.UserLogService,
		resetPasswordRepository: params.ResetPasswordRepository,
		securityHelper:          params.SecurityHelper,
		config:                  params.Config,
		transactor:              params.Transactor,
	}
}
