package mfa_recovery_code

import (
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_codes"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	service2 "github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaRecoveryCodeService struct {
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     mfa_recovery_codes.MfaRecoveryCodesHelper
	config                    Config
}

type Config struct {
	MfaRecoveryCodeCount int
}

type ServiceParams struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaRecoveryCodeHelper     mfa_recovery_codes.MfaRecoveryCodesHelper
	Config                    Config
}

func New(params ServiceParams) service2.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		mfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     params.MfaRecoveryCodeHelper,
		config:                    params.Config,
	}
}
