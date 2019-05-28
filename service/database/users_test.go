package database_test

import (
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authentication/service/database"
)

type (
	AnyTime      struct{}
	AnyByteArray struct {
	}
)

func TestUsers_Create_userEmail(t *testing.T) {

	mockDatabase, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO").
		ExpectExec().WithArgs("locked", "vendor", AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectPrepare("INSERT INTO").
		ExpectExec().
		WithArgs(1, "test@example.com", AnyByteArray{}, AnyByteArray{}, AnyTime{}, AnyTime{}, false, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	d := database.DB(mockDatabase, logrus.New())

	err = d.Users().Create(&database.AuthModel{
		Email:    "test@example.com",
		Password: []byte("pass"),
		Salt:     []byte("salt"),
	}, &database.UsersModel{
		Status_account: "locked",
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

	mockDatabase, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO").
		ExpectExec().WithArgs("active", "vendor", AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(30, 1))

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(30).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectPrepare("INSERT INTO").
		ExpectExec().WithArgs(30, "3", "3333333333", AnyByteArray{}, AnyByteArray{}, AnyTime{}, AnyTime{}, false, false).WillReturnResult(sqlmock.NewResult(30, 1))

	mock.ExpectCommit()
	d := database.DB(mockDatabase, logrus.New())
	err = d.Users().Create(&database.AuthModel{
		PhoneCountryCode: "3",
		PhoneNumber:      "3333333333",
		Password:         []byte("users_create_phone_pass"),
		Salt:             []byte("users_create_phone_salt"),
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

func TestUser_Update_StatusAccount(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE users SET status_account = ?, updated=? WHERE id = ?").WithArgs("active", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Users().Update_StatusAccount(&database.UsersModel{
		Id:             1,
		Status_account: "active",
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Update_Type(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE users SET type=?, updated=? WHERE id = ? ").WithArgs("vendor", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Users().Update_Type(&database.UsersModel{
		Id:   1,
		Type: "vendor",
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_Update_MfaType(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE users SET mfa_type=? , updated=? WHERE id = ?").WithArgs("opt", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Users().Update_MfaType(&database.UsersModel{
		Id:       1,
		Mfa_type: "opt",
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func (s AnyByteArray) Match(v driver.Value) bool {
	_, ok := v.([]byte)
	return ok
}
