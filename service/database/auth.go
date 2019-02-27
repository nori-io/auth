package database

import (
	"database/sql"
	"errors"
)

type auth struct {
	db *sql.DB
}

func (a *auth) Create(model *AuthModel) error {
	_, err := a.db.Exec("INSERT INTO auth (user_id, phone, email, password, salt, created, updated, is_email_verified, is_phone_verified) VALUES(?,?,?,?,?,?,?,?,?)",
		model.UserId, model.Phone, model.Email, model.Password, model.Salt, model.Created, model.Updated, model.IsEmailVerified, model.IsPhoneVerified)
	return err
}

func (a *auth) Update(model *AuthModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth SET profile_user_id = ?, phone = ?, email = ?, password = ? salt = ? created =? WHERE id = ? ",
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
