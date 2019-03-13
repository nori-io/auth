package database_test

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/nori-io/auth/service/database"
	"github.com/nori-io/auth/service/database/sql_scripts"
)

func TestUsers_Create(t *testing.T) {

	type Users interface {
		Create(*database.AuthModel, *database.UsersModel) error
	}
	var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel

	var err error

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}


	tx, err := db.Begin()
log.Print(err)
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
		sql_scripts.CreateTableUsersMfaPhone)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaCode)
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

	defer db.Close()

	modelUsers = &database.UsersModel{
		Type:    "vendor",
		Created: time.Now(),
		Updated: time.Now(),
	}
	modelAuth = &database.AuthModel{
		Email:    "test@mail.ru",
		Password: "pass",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	testDb := database.DB(db, nil)
	// now we execute our method
	if err = testDb.Users().Create(modelAuth, modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}




}
