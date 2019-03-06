package database

import (
	"database/sql"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
)

type users struct {
	db  *sql.DB
	log *log.Logger
}

func (u *users) CreateAuth(modelAuth *AuthModel, modelUsers *UsersModel) error {
	var (
		lastIdNumber uint64
	)
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	_, execErr := tx.Exec("INSERT INTO users (status_account, type, created, updated, mfa_type) VALUES(?,?,?,?,?)",
		"active", modelUsers.Type, time.Now(), time.Now(), modelUsers.Mfa_type)
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}

	lastId, err := tx.Query("SELECT id FROM users WHERE id = (SELECT MAX(id) FROM users)")
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

	if (modelAuth.Phone == "") && (modelAuth.Email != "") {
		log.Println("Email add")
		_, execErr = tx.Exec("INSERT INTO auth (user_id,  email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)",
			lastIdNumber, modelAuth.Email, modelAuth.Password, modelAuth.Salt, time.Now(), time.Now(), false, false)
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}
	}

	if (modelAuth.Phone != "") && (modelAuth.Email == "") {
		log.Println("Phone add")
		_, execErr = tx.Exec("INSERT INTO auth (user_id, phone, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?)",
			lastIdNumber, modelAuth.Phone, modelAuth.Password, modelAuth.Salt, time.Now(), time.Now(), false, false)
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
