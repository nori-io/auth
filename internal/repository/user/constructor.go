package user

import (
	"github.com/jinzhu/gorm"

	interfaceUser "github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/repository/user/postgres"
)

func New(db *gorm.DB) interfaceUser.UserRepository {
	return &postgres.UserRepository{Db: db}
}
