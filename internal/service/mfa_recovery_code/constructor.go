package mfa_recovery_code

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type MfaRecoveryCodeService struct {
	authenticationLogService  service.AuthenticationLogService
	sessionService            service.SessionService
	userService               service.UserService
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	config                    config.Config
	transactor                transactor.Transactor
}

type Params struct {
	authenticationLogService  service.AuthenticationLogService
	sessionService            service.SessionService
	userService               service.UserService
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	Config                    config.Config
	Transactor                transactor.Transactor
}

func New(params Params) service.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		authenticationLogService:  params.authenticationLogService,
		userService:               params.userService,
		sessionService:            params.sessionService,
		mfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     params.MfaRecoveryCodeHelper,
		config:                    params.Config,
		transactor:                params.Transactor,
	}
}
