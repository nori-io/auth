package database_test

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/nori-io/auth/service/database"
	"github.com/nori-io/auth/service/database/sql_scripts"
)


type (
  AnyTime struct{}
)

func TestUsers_Create(t *testing.T) {
	type Database interface {
		Users() database.Users
		Auth() database.Auth
	}
	type Users interface {
		Create(*database.AuthModel, *database.UsersModel) error
	}

	var err error


	db1, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db1.Close()

	type users struct {
		db  *sql.DB
		log *log.Logger
	}

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
  userObject:=database.Users1{Db:db1, Log:nil}
	mock.ExpectBegin()

	mock.ExpectExec("INSERT INTO users").WithArgs("active","vendor",AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(2,10))


  if err = userObject.Create(modelAuth, modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}




}
func TestUsers_Create2(t *testing.T) {


	db, mock, err:= sqlmock.New()
	db.Query(sql_scripts.SetDatabaseSettings)
	db.Query(sql_scripts.SetDatabaseStricts)
	db.Query(sql_scripts.CreateTableUsers)

	defer db.Close()

	mock.ExpectExec("INSERT INTO users").
		WithArgs(1,"active", "vendor",AnyTime{},AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1,1))

	_, err = db.Exec("INSERT INTO users(id,status_account,type, created, updated) VALUES (?,?, ?,?,?)", 1,"active","vendor", time.Now(),time.Now())
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}
	/*lastId, err1:= db.Exec("SELECT id FROM users WHERE id = (SELECT MAX(id) FROM users)")
	if err1!= nil {
		log.Println(err1)
	}
	t.Log(lastId)
*/
	/*if lastId.Err() != nil {
		log.Println(err1)
	}

	defer lastId.Close()
	for lastId.Next() {
		var m database.UsersModel
		lastId.Scan(&m.Id)
		lastIdNumber:= m.Id
		t.Log("LastNumber is",lastIdNumber)
	}
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
*/


	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
