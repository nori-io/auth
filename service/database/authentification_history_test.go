package database_test

import (
	"runtime"
	"runtime/debug"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

func TestAuthenticationHistory_Create(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mock.ExpectExec("INSERT INTO authentification_history (user_id, signin, meta) VALUES(?,?,?)").
		WithArgs(1,  AnyTime{}, "").WillReturnResult(sqlmock.NewResult(1, 1))


	d := database.DB(mockDatabase, logrus.New())

	err = d.AuthenticationHistory().Create(&database.AuthenticationHistoryModel{
		UserId:1,
		SignIn:time.Now(),
		Meta:"",
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	clear(d)
	d = nil
	debug.SetGCPercent(1)
	runtime.GC()

}


func TestAuthenticationHistory_Update(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()
	mock.ExpectExec("UPDATE authentification_history SET user_id = ?, signin = ?, meta = ?, signout = ?  WHERE id = ? ").
		WithArgs(1,   AnyTime{},"",AnyTime{},1).WillReturnResult(sqlmock.NewResult(1, 0))


	d := database.DB(mockDatabase, logrus.New())

	err = d.AuthenticationHistory().Update(&database.AuthenticationHistoryModel{
		Id:1,
		UserId:1,
		SignIn:time.Now(),
		Meta:"",
		SignOut:time.Now(),
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}


}
