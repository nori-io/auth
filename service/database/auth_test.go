package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
)

func TestAuth_Update(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? ,salt = ? ,created =? WHERE id = ? ").WithArgs(1,1,).
		WillReturnResult(sqlmock.NewResult(1,0))


	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update(&database.AuthModel{
		Id:1,
		UserId:1,
		Email:    "test@mail.ru",
		Password: "pass",
		Salt:     "salt",
	})



	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

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

	nonEmptyRows := sqlmock.NewRows([]string{"id", "phone_country_code", "phone_number", "password"}).
		AddRow(1, "8", "9191501490", "pass")

	mock.ExpectQuery("SELECT id, phone_country_code, phone_number, password FROM auth WHERE phone_country_code = ? and phone_number=?  LIMIT 1").WithArgs("8", "9191501490").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByPhone("8", "9191501490")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
