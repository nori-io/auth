package reset_password

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/reset_password/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.ResetPasswordRepository {
	return &postgres.ResetPasswordRepository{Tx: tx}
}
