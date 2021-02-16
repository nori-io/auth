package postgres

import "github.com/jinzhu/gorm"

type AuthRepository struct {
	Db *gorm.DB
}

type Auth struct {
	Email    string `gorm:"column:email; type: VARCHAR(32)"`
	Password string `gorm:"column:email; type: VARCHAR(32)"`
}
