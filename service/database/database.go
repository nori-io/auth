package database

import (
	"context"
	"database/sql"

	"github.com/nori-io/nori-common/interfaces"
	log "github.com/sirupsen/logrus"

	"github.com/nori-io/authentication/service/database/sql_scripts"
)

type Database interface {
	Users() Users
	AuthenticationHistory() AuthenticationHistory
	Auth() Auth
	MfaRecoveryCodes() MfaRecoveryCodes

	CreateTables() error
	DropTables() error
}

type AuthenticationHistory interface {
	Create(*AuthenticationHistoryModel) error
	Update(*AuthenticationHistoryModel) error
}

type Users interface {
	Create(*AuthModel, *UsersModel) error

	Update_StatusAccount(modelUsers *UsersModel) error
	Update_Type(modelUsers *UsersModel) error
	Update_MfaType(modelUsers *UsersModel) error

	//	CreateProvider(*ProviderModel, *UsersModel) error
	//Update(*UsersModel) error
}

type Auth interface {
	Update_PhoneNumber_CountryCode(model *AuthModel) error
	Update_Email(model *AuthModel) error
	Update_Password_Salt(model *AuthModel) error
	Update_IsEmailVerified(model *AuthModel) error
	Update_IsPhoneVerified(model *AuthModel) error
	FindByEmail(email string) (model *AuthModel, err error)
	FindByPhone(phoneCountryCode, phoneNumber string) (model *AuthModel, err error)
}

type MfaRecoveryCodes interface {
	Create(modelMfaRecoveryCodes *MfaRecoveryCodesModel) ([]string, error)
	Delete(code string) error
}
type database struct {
	db                    *sql.DB
	users                 *user
	authenticationHistory *authenticationHistory
	auth                  *auth
	mfaRecoveryCodes      *mfaRecoveryCodes
}

func DB(db *sql.DB, logger interfaces.Logger) Database {

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
		mfaRecoveryCodes: &mfaRecoveryCodes{
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

func (db *database) MfaRecoveryCodes() MfaRecoveryCodes {
	return db.mfaRecoveryCodes
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
		sql_scripts.CreateTableAuthenticationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUserMfaRecoveryCodes)
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
		sql_scripts.DropTableAuthenticationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.DropTableUserMfaRecoveryCodes)
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
