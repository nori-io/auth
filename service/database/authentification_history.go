package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nori-io/nori-common/interfaces"
)

type authenticationHistory struct {
	db  *sql.DB
	log interfaces.Logger
}

func (a *authenticationHistory) Create(model *AuthenticationHistoryModel) error {

	_, err := a.db.Exec("INSERT INTO authentification_history (user_id, signin, meta) VALUES(?,?,?)",
		model.UserId, time.Now(), model.Meta)
	return err
}

func (a *authenticationHistory) Update(model *AuthenticationHistoryModel) error {

	if model.UserId == 0 {
		return errors.New("Empty model")
	}

	_, err := a.db.Exec("UPDATE authentification_history SET  signout = ?   WHERE user_id = ? ORDER BY id DESC LIMIT 1",
		model.SignOut, model.UserId)
	return err
}
