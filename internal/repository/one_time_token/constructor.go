package one_time_token

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/one_time_token/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.OneTimeTokenRepository {
	return &postgres.OneTimeTokenRepository{Tx: tx}
}
