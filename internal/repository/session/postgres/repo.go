package postgres

import "github.com/jinzhu/gorm"

type SessionRepository struct {
	Db *gorm.DB
}
