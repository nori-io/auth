package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type auth struct {
	db *sql.DB
}

func (a *auth) Create(model *AuthModel) error {
	var (
		lastIdNumber uint64
	)
	ctx := context.Background()

	tx, err := a.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	_, execErr := tx.Exec("INSERT INTO users (kind, status_id, type, created, updated, mfa_type) VALUES(?,?,?,?,?,?)",
		model.Users.Kind, model.Users.StatusId, model.Users.Type, model.Users.Created, model.Users.Updated, model.Users.Mfa_type)
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
		lastIdNumber, model.Phone, model.Email, model.Password, model.Salt, model.Created, model.Updated, model.IsEmailVerified, model.IsPhoneVerified)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatalf("Insert table 'auth' error", execErr.Error())
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Commit transaction error", err)
	}

	return nil

}

func (a *auth) Update(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? salt = ? created =? WHERE id = ? ",
		model.UserId, model.Id)
	return err
}

func (a *auth) FindByEmail(email string) (model *AuthModel, err error) {
	rows, err := a.db.Query("SELECT user_id, phone, email, password, salt, created, updated, is_email_verified, is_phone_verified FROM auth WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}

	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.Email)
		model.Id = m.Id
		model.Email = m.Email
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
