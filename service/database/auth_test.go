package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/authentication/service/database"
)

func TestAuth_Update_Email(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET email=? , updated=? WHERE id = ?").WithArgs("test@example.com", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_Email(&database.AuthModel{
		Id:    1,
		Email: "test@example.com",
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuth_Update_PhoneNumber_CountryCode(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET  phone_country_code =? , phone_number = ?, updated=? WHERE id = ?").WithArgs("1", "234567890", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_PhoneNumber_CountryCode(&database.AuthModel{
		Id:               1,
		PhoneCountryCode: "1",
		PhoneNumber:      "234567890",
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuth_Update_Password_Salt(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET password=? , salt=? , updated=? WHERE id = ?").WithArgs(AnyByteArray{}, AnyByteArray{}, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_Password_Salt(&database.AuthModel{
		Id:       1,
		Password: []byte("pass"),
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuth_UpdateIsEmailVerified(t *testing.T) {

	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET is_email_verified=? , updated=? WHERE id = ? ").WithArgs(true, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_IsEmailVerified(&database.AuthModel{
		Id:              1,
		IsEmailVerified: true,
	})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuth_UpdateIsPhoneVerified(t *testing.T) {
	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("UPDATE auth SET is_phone_verified=? , updated=? WHERE id = ? ").WithArgs(true, AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	d := database.DB(mockDatabase, logrus.New())

	err = d.Auth().Update_IsPhoneVerified(&database.AuthModel{
		Id:              1,
		IsPhoneVerified: true,
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
		AddRow(2, "auth_find_by_email_test@example.com", "auth_find_by_email_pass")

	mock.ExpectQuery("SELECT id, email,password, salt FROM auth WHERE email = ? LIMIT 1").WithArgs("auth_find_by_email_test@example.com").WillReturnRows(nonEmptyRows)

	d := database.DB(mockDatabase, logrus.New())

	_, err = d.Auth().FindByEmail("auth_find_by_email_test@example.com")

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
