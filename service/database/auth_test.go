package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authentication/service/database"
)

func TestAuth_Update(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? ,salt = ? ,created =? WHERE id = ? ").WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 0))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_Email(&database.AuthModel{
		Id:       1,
		UserId:   1,
		Email:    "auth_update_@mail.ru",
		Password: []byte("auth_update_pass"),
		Salt:     []byte("auth_update_salt"),
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
		AddRow(2, "auth_find_by_email_test@mail.ru", "auth_find_by_email_pass")

	mock.ExpectQuery("SELECT id, email,password, salt FROM auth WHERE email = ? LIMIT 1").WithArgs("auth_find_by_email_test@mail.ru").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByEmail("auth_find_by_email_test@mail.ru")

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
		AddRow(3, "1", "1111111111", "auth_find_by_phone_pass")

	mock.ExpectQuery("SELECT id, phone_country_code, phone_number, password FROM auth WHERE (phone_country_code+phone_number)=?  LIMIT 1").WithArgs("11111111111").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByPhone("1", "1111111111")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
