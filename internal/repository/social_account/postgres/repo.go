package postgres

import "github.com/jinzhu/gorm"

type SocialAccountRepository struct {
	Db *gorm.DB
}
