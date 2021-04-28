package reset_password

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type ResetPasswordService struct {
	resetPasswordRepository repository.ResetPasswordRepository
	config                  config.Config
	userService             service.UserService
	securityHelper          security.SecurityHelper
}

type Params struct {
	ResetPasswordRepository repository.ResetPasswordRepository
	Config                  config.Config
	userService             service.UserService
	securityHelper          security.SecurityHelper
}

func New(params Params) service.ResetPasswordService {
	return &ResetPasswordService{
		resetPasswordRepository: params.ResetPasswordRepository,
		config:                  params.Config,
		userService:             params.userService,
		securityHelper:          params.securityHelper,
	}
}
