package database_test

import (
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

type (
	AnyTime struct{}
)

func TestUsers_Create_userEmail(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users (status_account, type, created, updated) VALUES(?,?,?,?)").
		WithArgs("active", "vendor", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(20).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)").
		WithArgs(20, "users_create_email_test@mail.ru", "users_create_email_pass", "users_create_email_salt", AnyTime{}, AnyTime{}, false, false).WillReturnResult(sqlmock.NewResult(20, 1))
	mock.ExpectCommit()

	d := database.DB(mockDatabase, logrus.New())

	err = d.Users().Create(&database.AuthModel{
		Email:    "users_create_email_test@mail.ru",
		Password: "users_create_email_pass",
		Salt:     "users_create_email_salt",
	}, &database.UsersModel{
		Status_account: "active",
		Type:           "vendor",
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestUsers_Create_userPhone(t *testing.T) {

	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.ExpectClose()
	defer mockDatabase.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users (status_account, type, created, updated,mfa_type) VALUES(?,?,?,?,?)").
		WithArgs("active", "vendor", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(30, 1))
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(30).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO auth (user_id, phone_country_code, phone_number, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)").
		WithArgs(30, "3", "3333333333", "users_create_phone_pass", "users_create_phone_salt", AnyTime{}, AnyTime{}, false, false).WillReturnResult(sqlmock.NewResult(30, 1))

	mock.ExpectCommit()
	d := database.DB(mockDatabase, logrus.New())
	err = d.Users().Create(&database.AuthModel{
		PhoneCountryCode: "3",
		PhoneNumber:      "3333333333",
		Password:         "users_create_phone_pass",
		Salt:             "users_create_phone_salt",
	}, &database.UsersModel{
		Status_account: "active",
		Type:           "vendor",
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
