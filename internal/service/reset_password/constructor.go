package reset_password

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type ResetPasswordService struct {
	authenticationLogService service.AuthenticationLogService
	resetPasswordRepository  repository.ResetPasswordRepository
	config                   config.Config
	userService              service.UserService
	securityHelper           security.SecurityHelper
	transactor               transactor.Transactor
}

type Params struct {
	authenticationLogService service.AuthenticationLogService
	ResetPasswordRepository  repository.ResetPasswordRepository
	Config                   config.Config
	userService              service.UserService
	securityHelper           security.SecurityHelper
	transactor               transactor.Transactor
}

func New(params Params) service.ResetPasswordService {
	return &ResetPasswordService{
		authenticationLogService: params.authenticationLogService,
		resetPasswordRepository:  params.ResetPasswordRepository,
		config:                   params.Config,
		userService:              params.userService,
		securityHelper:           params.securityHelper,
		transactor:               params.transactor,
	}
}
