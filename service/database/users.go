package database

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
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

	stmt, err := tx.Prepare("INSERT INTO users (status_account, type, created, updated) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, execErr := stmt.Exec("active", modelUsers.Type, time.Now(), time.Now())
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

		salt, err := Randbytes(65)
		if err != nil {
			return err
		}

		password, err := HashPassword([]byte(modelAuth.Password), salt)
		if err != nil {
			return err
		}

		encodedPassword := base64.StdEncoding.EncodeToString(password)
		encodedSalt := base64.StdEncoding.EncodeToString(salt)

		stmt, err := tx.Prepare("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}

		_, execErr := stmt.Exec(lastIdNumber, modelAuth.Email, encodedPassword, encodedSalt, time.Now(), time.Now(), false, false)
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}
	}

	if (modelAuth.PhoneCountryCode+modelAuth.PhoneNumber != "") && (modelAuth.Email == "") {
		stmt, err := tx.Prepare("INSERT INTO auth (user_id, phone_country_code, phone_number, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		_, execErr = stmt.Exec(lastIdNumber, modelAuth.PhoneCountryCode, modelAuth.PhoneNumber, modelAuth.Password, modelAuth.Salt, time.Now(), time.Now(), false, false)
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

func (u *user) Update(modelUsers *UsersModel) error {
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	if modelUsers.Id == 0 {
		return errors.New("Empty model")
	}
	_, err = tx.Exec("UPDATE users SET status_account = ?, updated = ?, mfa_type = ?  WHERE id = ? ",
		modelUsers.Status_account, time.Now(), modelUsers.Mfa_type)
	return err
	return nil
}
