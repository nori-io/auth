package transactor_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mocks "github.com/nori-io/common/v4/pkg/domain/mocks/registry"
	config2 "github.com/nori-plugins/authentication/internal/config"

	"github.com/nori-plugins/authentication/internal/domain/service"

	userSrv "github.com/nori-plugins/authentication/internal/service/user"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/logger"
	"github.com/nori-plugins/authentication/internal/domain/entity"
	userRepository "github.com/nori-plugins/authentication/internal/repository/user"
	"github.com/nori-plugins/authentication/pkg/transactor"
	"github.com/stretchr/testify/require"
)

func TestTxManager_Transact(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	user := &entity.User{
		Status:                 users_status.Active,
		UserType:               1,
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
		UpdatedAt:              time.Now(),
	}
	sqlInsertString := `INSERT INTO "users" ` +
		`("status","user_type","mfa_type","phone_country_code","phone_number","email","password","salt","hash_algorithm","is_email_verified","is_phone_verified","email_activation_code ","email_activation_code_ttl","created_at","updated_at") ` +
		`VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING "users"."id"`

	mock.ExpectQuery(`SELECT * FROM "users"  WHERE (email=$1) ORDER BY "users"."id" ASC LIMIT 1`).
		WithArgs("1").WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsertString).
		WithArgs(user.Status, 1, 1, "1", "1", "1", "1", "1", 1, false, false, "1", time.Now(), time.Now(), time.Now()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT * FROM "users"  WHERE "users"."id" = $1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "status", "user_type", "mfa_type",
			"phone_country_code", "phone_number", "email",
			"password", "salt", "hash_algorithm",
			"is_email_verified", "is_phone_verified", "email_activation_code",
			"email_activation_code_ttl", "created_at", "updated_at",
		}).AddRow(1, user.Status, 1, 1, "1", "1", "1", "1", "1", 1, false, false, "1", time.Now(), time.Now(), time.Now()))

	gdb, err := gorm.Open("postgres", db)

	txParams := transactor.Params{
		Db:  gdb,
		Log: logger.L(),
	}
	tx := transactor.New(txParams)

	r := userRepository.New(tx)

	ctrl := &gomock.Controller{T: t}
	conf := mocks.NewMockConfig(ctrl)

	config := &config2.Config{
		PasswordBcryptCost: conf.Int("10", "10"),
	}

	s := userSrv.New(userSrv.Params{
		UserRepository: r,
		Transactor:     tx,
		Сonfig:         *config,
	})

	user, err = s.Create(context.Background(), service.UserCreateData{
		Email:    "1",
		Password: "1",
	})
	if err != nil {
		require.NoError(t, err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		require.NoError(t, err)
	}
}
