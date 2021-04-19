package mfa_recovery_code

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.MfaRecoveryCodeRepository {
	return &postgres.MfaRecoveryCodeRepository{Tx: tx}
}
