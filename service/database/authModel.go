package database

import (
	"database/sql"
	"errors"
)

type auth struct {
	db *sql.DB
}

func (a *auth) Create(model *AuthModel) error {


	_, err1 := a.db.Exec("INSERT INTO users (kind, status_id, type, created, updated, mfa_type) VALUES(?,?,?,?,?,?)",
		model.Kind_Users, model.StatusId_Users, model.Type_Users, model.Created_Users, model.Updated_Users, model.Mfa_type_Users)
	return err1

	_, err2 := a.db.Exec("INSERT INTO auth (user_id, phone, email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)",
		model.Id_Auth, model.Phone_Auth, model.Email_Auth, model.Password_Auth, model.Salt_Auth, model.Created_Auth, model.Updated_Auth, model.IsEmailVerified_Auth, model.IsPhoneVerified_Auth)
	return err2
}

func (a *auth) Update(model *AuthModel) error {
	if model.Id_Auth == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? salt = ? created =? WHERE id = ? ",
		model.UserId_Auth, model.Id_Auth)
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
		rows.Scan(&m.Id_Auth, &m.Email_Auth)
		model.Id_Auth = m.Id_Auth
		model.Email_Auth = m.Email_Auth
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
