package database

import (
	"database/sql"
	"errors"
)

type users struct {
	db *sql.DB
}

func (u *users) Create(model *UsersModel) error {
	_, err := u.db.Exec("INSERT INTO users (profile_type_id, status_id, kind, created, updated, email) VALUES(?,?,?,?,?,?)",
		model.ProfileTypeId, model.StatusId, model.Kind, model.Created, model.Updated, model.Email)
	return err
}

func (u *users) Update(model *UsersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := u.db.Exec("UPDATE users SET profile_type_id = ?, status_id = ?, kind = ?, created = ? updated = ? email =? WHERE id = ? ",
		model.ProfileTypeId, model.StatusId, model.Kind, model.Created, model.Updated, model.Email, model.Id)
	return err
}

