package database

import (
	"database/sql"

	"errors"
)

type user struct {
	db *sql.DB
}

func (u *user) Create(model *UserModel) error {
	_, err := u.db.Exec("INSERT INTO users (name,email,salt,password) VALUES(?,?,?,?)",
		model.Name, model.Email, model.Salt, model.Password)
	return err
}

func (u *user) Update(model *UserModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := u.db.Exec("UPDATE users SET name = ?, email = ?, salt = ?, password = ? WHERE id = ? ",
		model.Name, model.Email, model.Salt, model.Password, model.Id)
	return err
}

func (u *user) FindById(id string) (model *UserModel, err error) {
	rows, err := u.db.Query("SELECT id, name, email, salt, password FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	model = &UserModel{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&model.Id, &model.Name, &model.Email, &model.Salt, &model.Password)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}

func (u *user) FindByEmail(email string) (model *UserModel, err error) {
	rows, err := u.db.Query("SELECT id, name, email, salt, password FROM users WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	model = &UserModel{}

	defer rows.Close()
	for rows.Next() {
		var m UserModel
		rows.Scan(&m.Id, &m.Name, &m.Email, &m.Salt, &m.Password)
		model.Id = m.Id
		model.Name = m.Name
		model.Email = m.Email
		model.Password = m.Password
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
