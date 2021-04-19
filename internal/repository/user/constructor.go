package user

import (
	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/repository/user/postgres"
)

func New(tx transactor.Transactor) repository.UserRepository {
	return &postgres.UserRepository{
		Tx: tx,
	}
}
