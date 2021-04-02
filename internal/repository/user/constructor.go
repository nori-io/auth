package user

import (
	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/repository/user/postgres"
)

//func New(db *gorm.DB) repository.UserRepository {
//	return &postgres.UserRepository{Db: db}
//}
func New(tx transactor.Transactor) repository.UserRepository {
	return &postgres.UserRepository{
		//Db: db,
		Tx: tx,
	}
}
