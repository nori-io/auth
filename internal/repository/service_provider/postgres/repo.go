package postgres

import "github.com/jinzhu/gorm"

type ServiceProviderRepository struct {
	Db *gorm.DB
}
