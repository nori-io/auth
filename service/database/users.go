package database

import (
	"context"
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"
)

type user struct {
	db  *sql.DB
	Log *log.Logger
}

func (u *user) Create(modelAuth *AuthModel, modelUsers *UsersModel) error {
	var (
		lastIdNumber uint64
	)
	ctx := context.Background()
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	_, execErr := tx.Exec("INSERT INTO users (status_account, type, created, updated) VALUES(?,?,?,?)",
		"active", modelUsers.Type, time.Now(), time.Now())
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}

	lastId, err := tx.Query("SELECT LAST_INSERT_ID()")
	if err != nil {
		return err
	}
	if lastId.Err() != nil {
		return err
	}
	defer lastId.Close()
	for lastId.Next() {
		var m AuthModel
		lastId.Scan(&m.Id)
		lastIdNumber = m.Id
	}
	if (modelAuth.PhoneCountryCode+modelAuth.PhoneNumber == "") && (modelAuth.Email != "") {
		_, execErr = tx.Exec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)",
			lastIdNumber, modelAuth.Email, modelAuth.Password, modelAuth.Salt, time.Now(), time.Now(), false, false)
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}
	}

	if (modelAuth.PhoneCountryCode+modelAuth.PhoneNumber != "") && (modelAuth.Email == "") {
		_, execErr = tx.Exec("INSERT INTO auth (user_id, phone_country_code, phone_number, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)",
			lastIdNumber, modelAuth.PhoneCountryCode, modelAuth.PhoneNumber, modelAuth.Password, modelAuth.Salt, time.Now(), time.Now(), false, false)
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}

	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}
