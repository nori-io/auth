package mfa_recovery_code

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type MfaRecoveryCodeService struct {
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	config                    config.Config
	db                        *gorm.DB
	transactor                transactor.Transactor
}

type Params struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	Config                    config.Config
	Db                        *gorm.DB
	Transactor                transactor.Transactor
}

func New(params Params) service.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		mfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     params.MfaRecoveryCodeHelper,
		config:                    params.Config,
		db:                        params.Db,
		transactor:                params.Transactor,
	}
}
