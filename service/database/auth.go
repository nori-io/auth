package database

import (
	"database/sql"
	"errors"
)

type auth struct {
	db *sql.DB
}

func (a *auth) Create(model *AuthModel) error {
	_, err := a.db.Exec("INSERT INTO auth (user_id, phone, email, password, salt, created, updated, isEmailVerified, IsPhoneVerified) VALUES(?,?,?,?,?,?,?,?)",
		model.UserId, model.Phone, model.Email, model.Password, model.Salt, model.Created, model.Updated, model.IsEmailVerified, model.IsPhoneVerified)
	return err
}

func (a *auth) Update(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? salt = ? created =? WHERE id = ? ",
		model.UserId, model.Id)
	return err
}
