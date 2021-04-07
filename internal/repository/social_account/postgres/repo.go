package postgres

import (
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SocialAccountRepository struct {
	Tx transactor.Transactor
}
