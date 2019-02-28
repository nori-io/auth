package database

import (
	"database/sql"
		"log"

	"golang.org/x/net/context"
)

type users struct {
	db *sql.DB
}

func (u *users) CreateAuth(modelAuth *AuthModel, modelUsers *UsersModel) error {
	var (
		lastIdNumber uint64
	)
	ctx := context.Background()

	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Exec("INSERT INTO users ( status_id, type, created, updated, mfa_type) VALUES(?,?,?,?,?)",
		modelUsers.StatusId, modelUsers.Type, modelUsers.Created, modelUsers.Updated, modelUsers.Mfa_type)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatalf("Insert table 'users' error", execErr)
	}

	lastId, err := tx.Query("SELECT id FROM users WHERE id = (SELECT MAX(id) FROM users)")
	if err != nil {
		log.Fatalf("Select table 'users' error ", err)
	}

	if lastId.Err() != nil {
		log.Fatalf("Taking lastId error ", err)
	}

	defer lastId.Close()
	for lastId.Next() {
		var m AuthModel
		lastId.Scan(&m.Id)
		lastIdNumber = m.Id
	}

	_, execErr = tx.Exec("INSERT INTO auth (user_id, phone, email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)",
		lastIdNumber, modelAuth.Phone, modelAuth.Email, modelAuth.Password, modelAuth.Salt, modelAuth.Created, modelAuth.Updated, modelAuth.IsEmailVerified, modelAuth.IsPhoneVerified)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatalf("Insert table 'auth' error", execErr.Error())
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Commit transaction error", err)
	}

	return nil

}