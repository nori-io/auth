package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nori-io/nori-common/logger"
	
)

type authenticationHistory struct {
	db  *sql.DB
	log logger.Writer
}

func (a *authenticationHistory) Create(model *AuthenticationHistoryModel) error {

	_, err := a.db.Exec("INSERT INTO authentication_history (user_id, signin, meta) VALUES(?,?,?)",
		model.UserId, time.Now(), model.Meta)
	return err
}

func (a *authenticationHistory) Update(model *AuthenticationHistoryModel) error {

	if model.UserId == 0 {
		return errors.New("Empty model")
	}

	_, err := a.db.Exec("UPDATE authentication_history SET  signout = ?   WHERE user_id = ? ORDER BY id DESC LIMIT 1",
		model.SignOut, model.UserId)
	return err
}
