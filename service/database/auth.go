package database

import (
	"database/sql"
	"errors"

	"github.com/nori-io/nori-common/interfaces"
)

type auth struct {
	db  *sql.DB
	log interfaces.Logger
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

	rows, err := a.db.Query("SELECT id, email,password,salt FROM auth WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}

	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.Email, &m.Password, &m.Salt)
		model.Id = m.Id

		model.Email = m.Email
		model.Password = m.Password
		model.Salt = m.Salt
	}

	if rows.Err() != nil {
		return nil, errors.New("User not found")
	}

	return model, nil
}

func (a *auth) FindByPhone(phoneCountryCode, phoneNumber string) (model *AuthModel, err error) {
	rows, err := a.db.Query("SELECT id, phone_country_code, phone_number, password,salt FROM auth WHERE concat(phone_country_code,phone_number)=?  LIMIT 1", phoneCountryCode+phoneNumber)
	if err != nil {
		return nil, err
	}
	model = &AuthModel{}
	defer rows.Close()
	for rows.Next() {
		var m AuthModel
		rows.Scan(&m.Id, &m.PhoneCountryCode, &m.PhoneNumber, &m.Password, &m.Salt)
		model.Id = m.Id
		model.PhoneCountryCode = m.PhoneCountryCode
		model.PhoneNumber = m.PhoneNumber
		model.Salt = m.Salt
		model.Password = m.Password
	}

	if rows.Err() != nil {
		return nil, errors.New("User not found")
	}

	return model, nil
}
