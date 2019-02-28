package database

import (
	"database/sql"
	"errors"
)

type users struct {
	db *sql.DB
}

func (u *users) Create(model *UsersModel) error {
	_, err := u.db.Exec("INSERT INTO users (kind, status_id, type, created, updated, mfa_type) VALUES(?,?,?,?,?,?)",
		model.Kind, model.StatusId, model.Type, model.Created, model.Updated, model.Mfa_type)
	return err
}

func (u *users) Update(model *UsersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := u.db.Exec("UPDATE users SET kind = ?, status_id = ?, type = ?, created = ?, updated = ?, mfa_type =? WHERE id = ? ",
		model.Kind, model.StatusId, model.Type, model.Created, model.Updated, model.Mfa_type, model.Id)
	return err
}
