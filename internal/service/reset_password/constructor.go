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
	userService              service.UserService
	resetPasswordRepository  repository.ResetPasswordRepository
	securityHelper           security.SecurityHelper
	config                   config.Config
	transactor               transactor.Transactor
}

type Params struct {
	AuthenticationLogService service.AuthenticationLogService
	UserService              service.UserService
	ResetPasswordRepository  repository.ResetPasswordRepository
	SecurityHelper           security.SecurityHelper
	Config                   config.Config
	Transactor               transactor.Transactor
}

func New(params Params) service.ResetPasswordService {
	return &ResetPasswordService{
		authenticationLogService: params.AuthenticationLogService,
		userService:              params.UserService,
		resetPasswordRepository:  params.ResetPasswordRepository,
		securityHelper:           params.SecurityHelper,
		config:                   params.Config,
		transactor:               params.Transactor,
	}
}
