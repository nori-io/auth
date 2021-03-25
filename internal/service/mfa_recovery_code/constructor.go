package mfa_recovery_code

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type MfaRecoveryCodeService struct {
	mfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	mfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	config                    config.Config
	session                   session.Session
	db                        *gorm.DB
}

type Params struct {
	MfaRecoveryCodeRepository repository.MfaRecoveryCodeRepository
	MfaRecoveryCodeHelper     mfa_recovery_code.MfaRecoveryCodesHelper
	Config                    config.Config
	Session                   session.Session
	Db                        *gorm.DB
}

func New(params Params) service.MfaRecoveryCodeService {
	return &MfaRecoveryCodeService{
		mfaRecoveryCodeRepository: params.MfaRecoveryCodeRepository,
		mfaRecoveryCodeHelper:     params.MfaRecoveryCodeHelper,
		config:                    params.Config,
		session:                   params.Session,
		db:                        params.Db,
	}
}
