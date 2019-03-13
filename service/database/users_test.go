package database_test

import (

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nori-io/auth/service/database"
)


func TestUsers_Create(t *testing.T) {

	type Users interface {
		Create(*database.AuthModel, *database.UsersModel) error
}
	var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel

	var err error


	sqlmock, mock, err:= sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	testDb:= database.DB(sqlmock,nil)






   defer sqlmock.Close()

    mock.ExpectBegin()
	/*mock.ExpectExec("INSERT INTO users (status_account, type, created, updated)").
		WithArgs(  "active", "vendor", time.Now(), time.Now()).WillReturnResult(nil)

	mock.ExpectExec("INSERT INTO users (status_account, type, created, updated)").
		WithArgs("").WillReturnResult(nil)

	mock.ExpectCommit()*/


	modelUsers = &database.UsersModel{
		Status_account:"active",
		Type:"vendor",
		Created:time.Now(),
		Updated:time.Now(),
	}
	modelAuth = &database.AuthModel{
	Email:"test@mail.ru",
	Password:"pass",
	Created:time.Now(),
	Updated:time.Now(),
	}


	// now we execute our method
	if err = testDb.Users().Create(modelAuth, modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	/*if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}*/
}