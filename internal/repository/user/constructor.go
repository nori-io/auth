package user

import (
	"github.com/jinzhu/gorm"

	interfaceUser "github.com/nori-io/authentication/internal/domain/repository"

	"github.com/nori-io/authentication/internal/repository/user/postgres"
)

func New(db *gorm.DB) interfaceUser.UserRepository {
	return &postgres.UserRepository{Db: db}
}
