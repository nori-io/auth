package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nori-io/nori-common/interfaces"
)

type auth struct {
	db  *sql.DB
	log interfaces.Logger
}

func (a *auth) Update_PhoneNumber_CountryCode(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET phone_number = ?, phone_country_code =? WHERE id = ? ",
		model.PhoneNumber, model.PhoneCountryCode, model.Id)

	if err == nil {
		a.Update_Updated(model)
	}
	return err
}
func (a *auth) Update_Email(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET email=? WHERE id = ? ",
		model.Email, model.Id)
	if err == nil {
		a.Update_Updated(model)
	}
	return err
}

func (a *auth) Update_Password_Salt(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}

	salt, err := createSalt()
	if err != nil {
		return err
	}

	password, err := Hash([]byte(model.Password), salt)
	if err != nil {
		return err
	}
	_, err = a.db.Exec("UPDATE auth SET password=? , salt=?  WHERE id = ? ",
		password, salt, model.Id)
	if err == nil {
		a.Update_Updated(model)
	}
	return err
}

func (a *auth) UpdateIsEmailVerified(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET is_email_verified=? WHERE id = ? ",
		model.IsEmailVerified, model.Id)
	if err == nil {
		a.Update_Updated(model)
	}
	return err
}

func (a *auth) UpdateIsPhoneVerified(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET is_phone_verified=? WHERE id = ? ",
		model.IsPhoneVerified, model.Id)
	if err == nil {
		a.Update_Updated(model)
	}
	return err
}

func (a *auth) Update_Updated(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET updated=? WHERE id = ? ",
		time.Now(), model.Id)

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
		return nil, rows.Err()
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
		return nil, rows.Err()
	}

	return model, nil
}
