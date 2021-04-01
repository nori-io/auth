package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationService struct {
	config                      config.Config
	authenticationLogRepository repository.AuthenticationLogRepository
	userRepository              repository.UserRepository
	sessionRepository           repository.SessionRepository
	mfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	db                          *gorm.DB

	userService              service.UserService
	authenticationLogService service.AuthenticationLogService
}

type Params struct {
	Config                      config.Config
	AuthenticationLogRepository repository.AuthenticationLogRepository
	MfaRecoveryCodeRepository   repository.MfaRecoveryCodeRepository
	SessionRepository           repository.SessionRepository
	UserRepository              repository.UserRepository
	DB                          *gorm.DB
	UserService                 service.UserService
	AuthenticationLogService    service.AuthenticationLogService
}

func New(params Params) service.AuthenticationService {
	return &AuthenticationService{
		config:                      params.Config,
		authenticationLogRepository: params.AuthenticationLogRepository,
		userRepository:              params.UserRepository,
		sessionRepository:           params.SessionRepository,
		mfaRecoveryCodeRepository:   params.MfaRecoveryCodeRepository,
		userService:                 params.UserService,
		authenticationLogService:    params.AuthenticationLogService,
		db:                          params.DB,
	}
}
