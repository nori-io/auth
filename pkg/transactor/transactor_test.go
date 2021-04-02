package transactor_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/logger"
	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/repository/user"
	"github.com/nori-plugins/authentication/pkg/transactor"
	"github.com/stretchr/testify/require"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestTxManager_Transact(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" ("status") VALUES ($1) RETURNING "users"."id"`).
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	mock.ExpectQuery(`SELECT * FROM "users"  WHERE "users"."id" = $1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	/*rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		RowError(1, fmt.Errorf("row error"))
	mock.ExpectQuery("SELECT LAST_INSERT_ID()").WillReturnRows(rows)

	mock.ExpectPrepare("INSERT INTO users").
		ExpectExec().
		WithArgs(1, 1, 1, 1, "1", "1", "1", "1", "1", 1, false, false, "1", AnyTime{}, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	*/
	//mock.ExpectCommit()
	gdb, err := gorm.Open("postgres", db)

	txParams := transactor.Params{
		Db:  gdb,
		Log: logger.L(),
	}
	tx := transactor.New(txParams)

	r := user.New(tx)

	err = r.Create(context.Background(), &entity.User{
		Status: 1,
		/*		UserType:               1,
				MfaType:                1,
				PhoneCountryCode:       "1",
				PhoneNumber:            "1",
				Email:                  "1",
				Password:               "1",
				Salt:                   "1",
				HashAlgorithm:          1,
				IsEmailVerified:        false,
				IsPhoneVerified:        false,
				EmailActivationCode:    "1",
				EmailActivationCodeTTL: time.Now(),
				CreatedAt:              time.Now(),
				UpdatedAt:              time.Now(),*/
	})
	if err != nil {
		require.NoError(t, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		require.NoError(t, err)
	}
}
