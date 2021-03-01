package user

import (
	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/internal/domain/repository"

	"github.com/nori-plugins/authentication/internal/repository/user/postgres"
)

func New(db *gorm.DB) repository.UserRepository {
	return &postgres.UserRepository{Db: db}
}
