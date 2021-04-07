package session

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/session/postgres"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

func New(tx transactor.Transactor) repository.SessionRepository {
	return &postgres.SessionRepository{Tx: tx}
}
