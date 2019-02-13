
package database

import (
	"database/sql"
	"errors"
)

type authenticationHistory struct {
	db *sql.DB
}


func (a *authenticationHistory) Create(model *AuthenticationHistoryModel) error {
	_, err := a.db.Exec("INSERT INTO authentication_history (user_id, logged_in, meta, logged_out, secret) VALUES(?,?,?,?,?)",
		model.UserId, model.LoggedIn, model.Meta, model.LoggedOut, model.Secret)
	return err
}

func (a *authenticationHistory) Update(model *AuthenticationHistoryModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE authentication_history SET user_id = ?, logged_in = ?, meta = ?, logged_out = ? secret = ?  WHERE id = ? ",
		model.UserId, model.LoggedIn, model.Meta, model.LoggedOut, model.Secret, model.Id)
	return err
}

