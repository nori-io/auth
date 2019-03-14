package database_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Selvatico/go-mocket"

	"github.com/nori-io/auth/service/database"
	"github.com/nori-io/auth/service/database/sql_scripts"
)


type (
  AnyTime struct{}
)

func TestUsers_Create(t *testing.T) {

	type Users interface {
		Create(*database.AuthModel, *database.UsersModel) error
	}
    var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel

	var err error
	ctx := context.Background()




	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(sql_scripts.SetDatabaseSettings)
	mock.ExpectExec(sql_scripts.SetDatabaseStricts)
	mock.ExpectExec(sql_scripts.CreateTableUsers)
	mock.ExpectExec(sql_scripts.CreateTableAuth)
	mock.ExpectExec("INSERT INTO users").WithArgs("active","vendor",time.Now(), time.Now()).WillReturnResult(sqlmock.NewResult(2,10))

	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))




	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}


/*	_, execErr:= tx.Exec(
		sql_scripts.SetDatabaseStricts)
	if execErr != nil {
		_ = tx.Rollback()

		log.Fatal(execErr)

	}*/

/*	_, execErr= tx.Exec(
		sql_scripts.SetDatabaseSettings)
	if execErr != nil {

		_ = tx.Rollback()
		log.Print("err is",execErr)

		log.Fatal(execErr)

	}*/



	_, execErr:= tx.Exec(
		sql_scripts.CreateTableUsers)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuth)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthProviders)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthentificationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaPhone)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaCode)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaSecret)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	modelUsers = &database.UsersModel{
		Type:    "vendor",
		Created: time.Now(),
		Updated: time.Now(),
	}
	modelAuth = &database.AuthModel{
		Email:    "test@mail.ru",
		Password: "pass",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	testDb := database.DB(db, nil)
	// now we execute our method

  if err = testDb.Users().Create(modelAuth, modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	mock.ExpectCommit()



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
func TestUsers_Create3(t *testing.T) {

	 gomocket.Catcher.Register()
/*	db, _ := sql.Open(DriverName, "connection_string") // Could be any connection string
	DB = db
	commonReply := []map[string]interface{}{{"name": "FirstLast", "age": "30"}}

	t.Run("Simple SELECT caught by query", func(t *testing.T) {
		Catcher.Logging = true
		fr := Catcher.Reset().NewMock().WithQuery(`SELECT name, age FROM users WHERE`).WithReply(commonReply)
		t.Log("result", fr)
		result := GetUsers(DB)*/
		t.Log("result", result)
		if len(result) != 1 {
			t.Fatalf("Returned sets is not equal to 1. Received %d", len(result))
		}
		if result[0]["name"] != "FirstLast" {
			t.Errorf("Name is not equal. Got %v", result[0]["name"])
		}
	})	type Users interface {
		Create(*database.AuthModel, *database.UsersModel) error
	}
	var modelAuth *database.AuthModel
	var modelUsers *database.UsersModel

	var err error
	ctx := context.Background()




	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(sql_scripts.SetDatabaseSettings)
	mock.ExpectExec(sql_scripts.SetDatabaseStricts)
	mock.ExpectExec(sql_scripts.CreateTableUsers)
	mock.ExpectExec(sql_scripts.CreateTableAuth)
	mock.ExpectExec("INSERT INTO users").WithArgs("active","vendor",time.Now(), time.Now()).WillReturnResult(sqlmock.NewResult(2,10))

	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))




	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}


	/*	_, execErr:= tx.Exec(
			sql_scripts.SetDatabaseStricts)
		if execErr != nil {
			_ = tx.Rollback()

			log.Fatal(execErr)

		}*/

	/*	_, execErr= tx.Exec(
			sql_scripts.SetDatabaseSettings)
		if execErr != nil {

			_ = tx.Rollback()
			log.Print("err is",execErr)

			log.Fatal(execErr)

		}*/



	_, execErr:= tx.Exec(
		sql_scripts.CreateTableUsers)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)

	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuth)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthProviders)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableAuthentificationHistory)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaPhone)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaCode)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	_, execErr = tx.Exec(
		sql_scripts.CreateTableUsersMfaSecret)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	modelUsers = &database.UsersModel{
		Type:    "vendor",
		Created: time.Now(),
		Updated: time.Now(),
	}
	modelAuth = &database.AuthModel{
		Email:    "test@mail.ru",
		Password: "pass",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	testDb := database.DB(db, nil)
	// now we execute our method

	if err = testDb.Users().Create(modelAuth, modelUsers); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	mock.ExpectCommit()



}
