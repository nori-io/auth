package database

import (
	"database/sql"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type authenticationHistory struct {
	db  *sql.DB
	log *log.Logger
}

func (a *authenticationHistory) Create(model *AuthenticationHistoryModel) error {
	_, err := a.db.Exec("INSERT INTO authentication_history (user_id, logged_in, meta) VALUES(?,?,?)",
		model.UserId, time.Now(), model.Meta)
	return err
}

func (a *authenticationHistory) Update(model *AuthenticationHistoryModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE authentication_history SET user_id = ?, logged_in = ?, meta = ?, logged_out = ?  WHERE id = ? ",
		model.UserId, model.LoggedIn, model.Meta, model.LoggedOut, model.Id)
	return err
}
