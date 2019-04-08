package database_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authorization/service/database"
)

func TestAuthenticationHistory_Create(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mock.ExpectExec("INSERT INTO authentification_history (user_id, signin, meta) VALUES(?,?,?)").
		WithArgs(10, AnyTime{}, "").WillReturnResult(sqlmock.NewResult(10, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.AuthenticationHistory().Create(&database.AuthenticationHistoryModel{
		UserId: 10,
		SignIn: time.Now(),
		Meta:   "",
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuthenticationHistory_Update(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mock.ExpectExec("UPDATE authentification_history SET user_id = ?, signin = ?, meta = ?, signout = ?  WHERE id = ? ").
		WithArgs(11, AnyTime{}, "", AnyTime{}, 11).WillReturnResult(sqlmock.NewResult(11, 0))

	d := database.DB(mockDatabase, logrus.New())

	err = d.AuthenticationHistory().Update(&database.AuthenticationHistoryModel{
		Id:      11,
		UserId:  11,
		SignIn:  time.Now(),
		Meta:    "",
		SignOut: time.Now(),
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
