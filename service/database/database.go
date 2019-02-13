package database

import (
	"database/sql"
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
	Create(*UsersModel) error
	Update(*UsersModel) error
}

type Auth interface {
	Create(model *AuthModel) error
	Update(*AuthModel) error
}
type database struct {
	db                    *sql.DB
	users                 *users
	authenticationHistory *authenticationHistory
	auth                  *auth
}

var instance *database
var once sync.Once

// Create Database using singltone pattern
func DB(db *sql.DB) Database {
	once.Do(func() {
		instance = &database{
			db: db,
			users: &users{
				db: db,
			},
			authenticationHistory: &authenticationHistory{
				db: db,
			},
			auth: &auth{
				db: db,
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
