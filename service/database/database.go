package database

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/nori-io/authorization/service/database/sql_scripts"
)

type Database interface {
	Users() Users
	AuthenticationHistory() AuthenticationHistory
	Auth() Auth
	MfaCode() MfaCode

	CreateTables() error
	DropTables() error
}

type AuthenticationHistory interface {
	Create(*AuthenticationHistoryModel) error
	Update(*AuthenticationHistoryModel) error
}

type Users interface {
	Create(*AuthModel, *UsersModel) error

	Update(*UsersModel) error

	//	CreateProvider(*ProviderModel, *UsersModel) error
	//Update(*UsersModel) error
}

type Auth interface {
	Update(*AuthModel) error
	FindByEmail(email string) (model *AuthModel, err error)
	FindByPhone(phoneCountryCode, phoneNumber string) (model *AuthModel, err error)
}

type MfaCode interface {
	Create(modelMfaCode *MfaCodeModel) ([]string, error)
	Delete(code string) error
}
type database struct {
	db                    *sql.DB
	users                 *user
	authenticationHistory *authenticationHistory
	auth                  *auth
	mfaCode               *mfaCode
}

func DB(db *sql.DB, logger *log.Logger) Database {

	return &database{
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
		mfaCode: &mfaCode{
			db:  db,
			log: logger,
		},
	}

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

func (db *database) MfaCode() MfaCode {
	return db.mfaCode
}

func (db *database) CreateTables() error {

	tx, err := db.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Exec(
		sql_scripts.SetDatabaseSettings)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}

	_, execErr = tx.Exec(
		sql_scripts.SetDatabaseStricts)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsers)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuth)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthProviders)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthentificationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUserMfaCode)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaPhone)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaSecret)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return nil
}
func (db *database) DropTables() error {

	tx, err := db.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Exec(
		sql_scripts.DropTableAuth)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}

	_, execErr = tx.Exec(
		sql_scripts.DropTableAuthProviders)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}
	_, execErr = tx.Exec(
		sql_scripts.DropTableAuthentificationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.DropTableUserMfaCode)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.DropTableUserMfaPhone)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.DropTableUserMfaSecret)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.DropTableUsers)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return nil
}
