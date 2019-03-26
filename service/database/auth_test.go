package database_test

import (
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

	nonEmptyRows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "test@mail.ru", "pass")

	mock.ExpectQuery("SELECT id, email,password FROM auth WHERE email = ? LIMIT 1").WithArgs("test@mail.ru").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByEmail("test@mail.ru")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuth_FindByPhone(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	nonEmptyRows := sqlmock.NewRows([]string{"id", "phone_country_code", "password"}).
		AddRow(1, "test@mail.ru", "pass")

	mock.ExpectQuery("SELECT id, email,password FROM auth WHERE email = ? LIMIT 1").WithArgs("test@mail.ru").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByEmail("test@mail.ru")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}