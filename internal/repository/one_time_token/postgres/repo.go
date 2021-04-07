package postgres

import "github.com/nori-plugins/authentication/pkg/transactor"

type OneTimeTokenRepository struct {
	Tx transactor.Transactor
}
