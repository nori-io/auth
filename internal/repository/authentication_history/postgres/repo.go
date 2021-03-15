package postgres

import "github.com/jinzhu/gorm"

type AuthenticationHistoryRepository struct {
	Db *gorm.DB
}
