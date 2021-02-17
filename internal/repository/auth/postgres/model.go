package postgres

import "github.com/jinzhu/gorm"

type AuthRepository struct {
	Db *gorm.DB
}

type Auth struct {
	ID       uint64 `gorm:"column:id; PRIMARY_KEY; type:bigserial"`
	UserID   uint64 `gorm:"column:user_id; type: bigint"`
	Email    string `gorm:"column:email; type: VARCHAR(32)"`
	Password string `gorm:"column:email; type: VARCHAR(32)"`
}
