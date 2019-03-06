package database

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
)

type auth struct {
	db  *sql.DB
	log *log.Logger
}

func (a *auth) Update(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? ,salt = ? ,created =? WHERE id = ? ",
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

func (a *auth) FindByPhone(phoneNumber string) (model *AuthModel, err error) {
	rows, err := a.db.Query("SELECT user_id, phone, email, password, salt, created, updated, is_email_verified, is_phone_verified FROM auth WHERE phone = ? LIMIT 1", phoneNumber)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}

	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.Phone)
		model.Id = m.Id
		model.Phone = m.Phone
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
