package session

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/session/postgres"
)

func New(db *gorm.DB) repository.SessionRepository {
	return &postgres.SessionRepository{Db: db}
}
