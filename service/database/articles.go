package database

import (
	"database/sql"
	"errors"
)

type article struct {
	db *sql.DB
}

//Create NEW record
func (a *article) Create(model *ArticlesModel) error {
	_, err := a.db.Exec("INSERT INTO articles (title,body,state,meta_description,tags) VALUES(?,?,?,?,?)",
		model.Title, model.Body, model.State, model.MetaDescription, model.Tags)
	return err
}

//Update existing record
func (a *article) Update(model *ArticlesModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE articles SET title = ?, body = ?, state = ?, meta_description = ?, tags=? WHERE id = ? ",
		model.Title, model.Body, model.State, model.MetaDescription, model.Tags, model.Id)
	return err
}

//
func (a *article) FindById(id int64) (model *ArticlesModel, err error) {
	rows, err := a.db.Query("SELECT id, title,body,state,meta_description,tags FROM articles WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	model = &ArticlesModel{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&model.Id, &model.Title, &model.Body, &model.State, &model.MetaDescription, &model.Tags)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
