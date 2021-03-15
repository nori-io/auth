package postgres

import "github.com/jinzhu/gorm"

type OneTimeTokenRepository struct {
	Db *gorm.DB
}
