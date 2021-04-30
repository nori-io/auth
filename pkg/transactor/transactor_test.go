package transactor_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	securityHelper2 "github.com/nori-plugins/authentication/internal/helper/security"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"

	"github.com/nori-plugins/authentication/pkg/enum/hash_algorithm"

	"github.com/nori-plugins/authentication/pkg/enum/mfa_type"

	"github.com/nori-plugins/authentication/pkg/enum/users_type"

	"github.com/nori-io/common/v4/pkg/domain/config"

	config2 "github.com/nori-plugins/authentication/internal/config"

	"github.com/nori-plugins/authentication/internal/domain/service"

	authSrv "github.com/nori-plugins/authentication/internal/service/auth"
	userSrv "github.com/nori-plugins/authentication/internal/service/user"
	userLogSrv "github.com/nori-plugins/authentication/internal/service/user_log"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/logger"
	"github.com/nori-plugins/authentication/internal/domain/entity"
	userRepository "github.com/nori-plugins/authentication/internal/repository/user"
	userLogRepository "github.com/nori-plugins/authentication/internal/repository/user_log"

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

	user := &entity.User{
		Status:                 users_status.Active,
		UserType:               users_type.User,
		MfaType:                mfa_type.None,
		PhoneCountryCode:       "",
		PhoneNumber:            "",
		Email:                  "test@mail.ru",
		Salt:                   "",
		HashAlgorithm:          hash_algorithm.Bcrypt,
		IsEmailVerified:        false,
		IsPhoneVerified:        false,
		EmailActivationCode:    "",
		EmailActivationCodeTTL: time.Now(),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
	sqlInsertString := `INSERT INTO "users" ` +
		`("status","user_type","mfa_type","phone_country_code","phone_number","email","password","salt","hash_algorithm","is_email_verified","is_phone_verified","email_activation_code ","email_activation_code_ttl","created_at","updated_at") ` +
		`VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING "users"."id"`

	mock.ExpectQuery(`SELECT * FROM "users"  WHERE (email=$1) ORDER BY "users"."id" ASC LIMIT 1`).
		WithArgs(user.Email).WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectBegin()
	mock.ExpectQuery(sqlInsertString).
		WithArgs(user.Status, user.UserType, user.MfaType, user.PhoneCountryCode, user.PhoneNumber, user.Email, sqlmock.AnyArg(), user.Salt, user.HashAlgorithm, user.IsEmailVerified, user.IsPhoneVerified, user.EmailActivationCode, AnyTime{}, AnyTime{}, AnyTime{}).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	gdb, err := gorm.Open("postgres", db)

	txParams := transactor.Params{
		Db:     gdb,
		Logger: logger.L(),
	}
	tx := transactor.New(txParams)

	r := userRepository.New(tx)

	config := &config2.Config{
		PasswordBcryptCost: func() config.Int {
			return func() int {
				return 10
			}
		}(),
		EmailVerification: func() config.Bool {
			return func() bool {
				return false
			}
		}(),
	}

	securityHelper := securityHelper2.New(securityHelper2.Params{Config: *config})

	s := userSrv.New(userSrv.Params{
		UserLogService: nil,
		UserRepository: r,
		SecurityHelper: securityHelper,
		Transactor:     tx,
		Config:         *config,
	})

	user, err = s.Create(context.Background(), service.UserCreateData{
		Email:    user.Email,
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

func TestTxManager_TransactNested(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	user := &entity.User{
		Status:                 users_status.Active,
		UserType:               users_type.User,
		MfaType:                mfa_type.None,
		PhoneCountryCode:       "",
		PhoneNumber:            "",
		Email:                  "test@mail.ru",
		Salt:                   "",
		HashAlgorithm:          hash_algorithm.Bcrypt,
		IsEmailVerified:        false,
		IsPhoneVerified:        false,
		EmailActivationCode:    "",
		EmailActivationCodeTTL: time.Now(),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
	user_log := &entity.UserLog{
		UserID:    1,
		Action:    users_action.SignUp,
		Meta:      "",
		CreatedAt: time.Now(),
	}
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT * FROM "users"  WHERE (email=$1) ORDER BY "users"."id" ASC LIMIT 1`).
		WithArgs(user.Email).WillReturnError(gorm.ErrRecordNotFound)
	sqlUserInsert := `INSERT INTO "users" ` +
		`("status","user_type","mfa_type","phone_country_code","phone_number","email","password","salt","hash_algorithm","is_email_verified","is_phone_verified","email_activation_code ","email_activation_code_ttl","created_at","updated_at") ` +
		`VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING "users"."id"`
	mock.ExpectQuery(sqlUserInsert).
		WithArgs(user.Status, user.UserType, user.MfaType, user.PhoneCountryCode, user.PhoneNumber, user.Email, sqlmock.AnyArg(), user.Salt, user.HashAlgorithm, user.IsEmailVerified, user.IsPhoneVerified, user.EmailActivationCode, AnyTime{}, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	sqlAuthenticationLogInsert := `INSERT INTO "nori_authentication_user_log" ("user_id","action","session_id","meta","created_at") VALUES ($1,$2,$3,$4,$5) RETURNING "nori_authentication_user_log"."id"`
	mock.ExpectQuery(sqlAuthenticationLogInsert).
		WithArgs(user_log.UserID, user_log.Action, sqlmock.AnyArg(), user_log.Meta, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	gdb, err := gorm.Open("postgres", db)

	txParams := transactor.Params{
		Db:     gdb,
		Logger: logger.L(),
	}
	tx := transactor.New(txParams)

	repoUser := userRepository.New(tx)
	repoUserLog := userLogRepository.New(tx)
	config := &config2.Config{
		PasswordBcryptCost: func() config.Int {
			return func() int {
				return 10
			}
		}(),
		EmailVerification: func() config.Bool {
			return func() bool {
				return false
			}
		}(),
	}
	securityHelper := securityHelper2.New(securityHelper2.Params{Config: *config})

	srvUserLog := userLogSrv.New(userLogSrv.Params{
		UserLogRepository: repoUserLog,
		Transactor:        tx,
	})

	srvUser := userSrv.New(userSrv.Params{
		UserLogService: srvUserLog,
		UserRepository: repoUser,
		SecurityHelper: securityHelper,
		Transactor:     tx,
		Config:         *config,
	})

	srvAuth := authSrv.New(authSrv.Params{
		Config:         *config,
		UserService:    srvUser,
		UserLogService: srvUserLog,
		Transactor:     tx,
	})

	user, err = srvAuth.SignUp(context.Background(), service.SignUpData{
		Email:    user.Email,
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
