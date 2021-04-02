package transactor_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/logger"
	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/repository/user"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestTxManager_Transact(t *testing.T) {
	var db *sql.DB
	var err error
	var mock sqlmock.Sqlmock
	db, mock, err = sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").WithArgs(0).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gdb, err := gorm.Open("postgres", db)

	txParams := transactor.Params{
		Db:  gdb,
		Log: logger.L(),
	}
	tx := transactor.New(txParams)

	r := user.New(tx)

	err = r.Create(context.Background(), &entity.User{
		Status: 1,
	})
	if err != nil {
		t.Error(err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
