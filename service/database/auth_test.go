package database_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

func TestAuth_FindByEmail(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mockDatabase.Exec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)",
		1, "test@mail.ru", "pass", "salt", time.Now(), time.Now(), false, false)

	mock.ExpectBegin()
	nonEmptyRows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1,"test@mail.ru", "pass")

	mock.ExpectQuery("SELECT FROM auth WHERE email = ? LIMIT 1").WithArgs("test@mail.ru").WillReturnRows(nonEmptyRows)

	mock.ExpectCommit()

	d := database.DB(mockDatabase, logrus.New())

	model, err:= d.Auth().FindByEmail("test@mail.ru")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Model",model)
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
