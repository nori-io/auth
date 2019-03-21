package database_test

import (
	"database/sql/driver"
	"fmt"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"

	"github.com/nori-io/auth/service/database"
	//"github.com/nori-io/auth/service/database"
)

type (
	AnyTime struct{}
)

func TestUsers_Create_userCreate(t *testing.T) {

	mockDatabase, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabase.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users (status_account, type, created, updated) VALUES(?,?,?,?)").
		WithArgs("active", "vendor", AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)").
		WithArgs(1, "test@mail.ru", "pass", "salt", AnyTime{}, AnyTime{}, false, false).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	d := database.DB(mockDatabase, logrus.New())
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

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

/*func TestUsers_Create2(t *testing.T) {
	var err error


	mockDatabase, mock, err := sqlmock.New()
	if !assert.Nilf(t,
		err,
		"Error on trying to start up a stub database connection") {
		t.Fatal()
	}
	defer mockDatabase.Close()

	testDatabase:=database.DB(mockDatabase,nil)



	modelUsers:=&database.UsersModel{
		Type:    "vendor",
		Created:time.Now(),
		Updated:time.Now(),

	}
	modelAuth:= &database.AuthModel{
		Email:    "test@mail.ru",
		Password: "pass",
		Created:time.Now(),
		Updated:time.Now(),

	}

	t.Run("WantCommitError", func(t *testing.T) {
		// define sql behavior
		mock.ExpectBegin()

		testDatabase.Users().Create(modelAuth,modelUsers)

		defer func() {
			mock.ExpectCommit().
				WillReturnError(nil)


		}()

		mock.ExpectationsWereMet()

	})


	if err = testDatabase.Users().Create(modelAuth,modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}




}
*/

func ExampleRows_rowError() {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(0, "one").
		AddRow(1, "two").
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	rs, _ := db.Query("SELECT")
	defer rs.Close()
	fmt.Println(rs)
	for rs.Next() {
		var id int
		var title string
		rs.Scan(&id, &title)
		fmt.Println("scanned id:", id, "and title:", title)
	}

	if rs.Err() != nil {
		fmt.Println("got rows error:", rs.Err())
	}
	// Output: scanned id: 0 and title: one
	// got rows error: row error
}
