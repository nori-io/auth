package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"

	"sync"
)

type Database interface {
	Users() Users
	AuthenticationHistory() AuthenticationHistory
	Auth() Auth
}

type AuthenticationHistory interface {
	Create(*AuthenticationHistoryModel) error
	Update(*AuthenticationHistoryModel) error
}

type Users interface {
	Create(*AuthModel, *UsersModel) error
	//	CreateProvider(*ProviderModel, *UsersModel) error
	//Update(*UsersModel) error
}

type Auth interface {
	//	Update(*AuthModel) error
	FindByEmail(email string) (model *AuthModel, err error)
	FindByPhone(phone string) (model *AuthModel, err error)
}

type Provider interface {
	//	Update(*ProviderModel) error
	//FindBy...
}
type database struct {
	db                    *sql.DB
	users                 *user
	authenticationHistory *authenticationHistory
	auth                  *auth
}

var instance *database
var once sync.Once

// Create Database using singltone pattern
func DB(db *sql.DB, logger *log.Logger) Database {
	once.Do(func() {
		instance = &database{
			db: db,
			users: &user{
				db:  db,
				Log: logger,
			},
			authenticationHistory: &authenticationHistory{
				db:  db,
				log: logger,
			},
			auth: &auth{
				db:  db,
				log: logger,
			},
		}
	})
	return instance
}

func (db *database) Users() Users {
	return db.users
}

func (db *database) AuthenticationHistory() AuthenticationHistory {
	return db.authenticationHistory
}

func (db *database) Auth() Auth {
	return db.auth
}
