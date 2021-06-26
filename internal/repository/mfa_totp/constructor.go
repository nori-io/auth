package mfa_totp

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_totp/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.MfaTotpRepository {
	return &postgres.MfaTotpRepository{Tx: tx}
}
