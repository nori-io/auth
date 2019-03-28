package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

func TestMfaCode_Create(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()
	defer mock.ExpectClose()

	d := database.DB(mockDatabase, logrus.New())

	mock.ExpectBegin()
	for index := 1; index <= 10; index++ {

		mock.ExpectExec("INSERT INTO user_mfa_phone (user_id, code) VALUES(?,?)").
			WithArgs(index, AnyTime{}).WillReturnResult(sqlmock.NewResult(int64(index), 1))

	}
	for index := 1; index <= 10; index++ {

		err = d.MfaCode().Create(&database.MfaCodeModel{UserId: uint64(index)})

	}

	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
