package authentication_log

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/authentication_log/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.AuthenticationLogRepository {
	return &postgres.AuthenticationLogRepository{Tx: tx}
}
