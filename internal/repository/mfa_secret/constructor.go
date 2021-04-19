package mfa_secret

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_secret/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.MfaSecretRepository {
	return &postgres.MfaSecretRepository{Tx: tx}
}
