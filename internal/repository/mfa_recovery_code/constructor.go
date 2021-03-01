package mfa_recovery_code

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code/postgres"
)

func New(db *gorm.DB) repository.MfaRecoveryCodeRepository {
	return &postgres.MfaRecoveryCodeRepository{Db: db}
}
