package mfa_recovery_code

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type MfaRecoveryCodeService struct {
	sessionService            service.SessionService
	userService               service.UserService
	userLogService            service.UserLogService
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     helper.MfaRecoveryCodesHelper
	config                    config.Config
	transactor                transactor.Transactor
}

type Params struct {
	SessionService            service.SessionService
	UserService               service.UserService
	UserLogService            service.UserLogService
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaRecoveryCodeHelper     helper.MfaRecoveryCodesHelper
	Config                    config.Config
	Transactor                transactor.Transactor
}

func New(params Params) service.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		sessionService:            params.SessionService,
		userService:               params.UserService,
		userLogService:            params.UserLogService,
		mfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     params.MfaRecoveryCodeHelper,
		config:                    params.Config,
		transactor:                params.Transactor,
	}
}
