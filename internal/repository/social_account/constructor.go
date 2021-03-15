package social_account

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/repository/social_account/postgres"
)

func New(db *gorm.DB) repository.SocialAccountRepository {
	return &postgres.SocialAccountRepository{Db: db}
}
