package database_test

import (
	"fmt"
	"testing"

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

	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO users (status_account, type, created, updated,mfa_type) VALUES(?,?,?,?,?)").
		WithArgs("active", "vendor", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)").
		WithArgs(1, "test@mail.ru", "pass", "salt", AnyTime{}, AnyTime{}, false, false).WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
	mock.ExpectBegin()

	mock.ExpectPrepare("INSERT INTO users (status_account, type, created, updated,mfa_type) VALUES(?,?,?,?,?)").ExpectExec().
		WithArgs("active", "vendor", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectPrepare("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)").ExpectExec().
		WithArgs(1, "test@mail.ru", "pass", "salt", AnyTime{}, AnyTime{}, false, false).WillReturnResult((sqlmock.NewResult(1, 1)))


	nonEmptyRows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1,"test@mail.ru", "pass")

	mock.ExpectQuery("SELECT FROM auth WHERE email = ? LIMIT 1").WithArgs("test@mail.ru").WillReturnRows(nonEmptyRows)

	mock.ExpectCommit()
    d := database.DB(mockDatabase,logrus.New())


	err = d.Users().Create(&database.AuthModel{
		Email:    "test@mail.ru",
		Password: "pass",
		Salt:     "salt",
	}, &database.UsersModel{
		Status_account: "active",
		Type:           "vendor",
	})
	if err != nil {
		t.Error(err)
	}
	_,err = d.Auth().FindByEmail("test@mail.ru")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
