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
	rows, err := a.db.Query("SELECT id, email,password FROM auth WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}

	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.Email, &m.Password)
		model.Id = m.Id
		model.Email = m.Email
		model.Password = m.Password

	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}

func (a *auth) FindByPhone(phone string) (model *AuthModel, err error) {
	rows, err := a.db.Query("SELECT id, phone, password FROM auth WHERE phone = ? LIMIT 1", phone)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}

	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.PhoneCountryCode,m.PhoneNumber, &m.Password)
		model.Id = m.Id
		model.PhoneCountryCode = m.PhoneCountryCode
		model.PhoneNumber=m.PhoneNumber
		model.Password = m.Password
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
