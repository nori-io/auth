package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nori-io/nori-common/interfaces"
)

type user struct {
	db  *sql.DB
	Log interfaces.Logger
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
	if len(modelUsers.Mfa_type) == 0 {
		stmt, err := tx.Prepare("INSERT INTO users (status_account, type, created, updated) VALUES(?,?,?,?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, execErr := stmt.Exec(modelUsers.Status_account, modelUsers.Type, time.Now(), time.Now())
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}
	} else {

		stmt, err := tx.Prepare("INSERT INTO users (status_account, type, created, updated,mfa_type) VALUES(?,?,?,?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, execErr := stmt.Exec(modelUsers.Status_account, modelUsers.Type, time.Now(), time.Now(), modelUsers.Mfa_type)
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}

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

	salt, err := CreateSalt()
	if err != nil {
		return err
	}

	password, err := Hash([]byte(modelAuth.Password), salt)
	if err != nil {
		return err
	}
	if (len(modelAuth.PhoneCountryCode+modelAuth.PhoneNumber) == 0) && (len(modelAuth.Email) != 0) {

		stmt, errInsertAuth := tx.Prepare("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)")
		if errInsertAuth != nil {
			return errInsertAuth
		}

		_, execErr := stmt.Exec(lastIdNumber, modelAuth.Email, password, salt, time.Now(), time.Now(), false, false)
		if execErr != nil {

			_ = tx.Rollback()
			return execErr
		}
	}

	if (len(modelAuth.PhoneCountryCode+modelAuth.PhoneNumber) != 0) && (len(modelAuth.Email) == 0) {
		stmt, err := tx.Prepare("INSERT INTO auth (user_id, phone_country_code, phone_number, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		_, execErr := stmt.Exec(lastIdNumber, modelAuth.PhoneCountryCode, modelAuth.PhoneNumber, password, salt, time.Now(), time.Now(), false, false)
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

func (u *user) Update_StatusAccount(modelUsers *UsersModel) error {
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	if modelUsers.Id == 0 {
		return errors.New("Empty model")
	}
	_, err = tx.Exec("UPDATE users SET status_account = ?, updated=? WHERE id = ? ",
		modelUsers.Status_account, time.Now(), modelUsers.Id)

	return err
}

func (u *user) Update_Type(modelUsers *UsersModel) error {
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	if modelUsers.Id == 0 {
		return errors.New("Empty model")
	}
	_, err = tx.Exec("UPDATE users SET type=?, updated=? WHERE id = ? ",
		modelUsers.Type, time.Now(), modelUsers.Id)
	tx.Commit()

	return err
}

func (u *user) Update_MfaType(modelUsers *UsersModel) error {
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	if modelUsers.Id == 0 {
		return errors.New("Empty model")
	}
	if modelUsers.Mfa_type == "" {
		_, err = tx.Exec("UPDATE users SET mfa_type=? , updated=? WHERE id = ?",
			sql.NullString{}, time.Now(), modelUsers.Id)
	} else {

		_, err = tx.Exec("UPDATE users SET mfa_type=? , updated=? WHERE id = ?",
			modelUsers.Mfa_type, time.Now(), modelUsers.Id)
	}
	return err
}
