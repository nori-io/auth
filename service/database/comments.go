package database

import (
	"database/sql"
	"errors"
)

type comment struct {
	db *sql.DB
}

//Create NEW record
func (c *comment) Create(model *CommentsModel) error {
	_, err := c.db.Exec("INSERT INTO comments (parent_id,post_id,message,created,state,article_id_fk) VALUES(?,?,?,?,?,?)",
		model.ParentId, model.PostId, model.Message, model.Created, model.State, model.ArticleIdFk)
	return err
}

//Update existing record
func (c *comment) Update(model *CommentsModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := c.db.Exec("UPDATE comments SET parent_id = ?, post_id = ?, message = ?, created = ?, state=? article_id_fk=?  WHERE id = ? ",
		model.ParentId, model.PostId, model.Message, model.Created, model.State, model.ArticleIdFk, model.Id)
	return err
}

func (c *comment) FindById(id int64) (model *CommentsModel, err error) {
	rows, err := c.db.Query("SELECT id, parent_id, post_id, message, created, state, article_id_fk FROM comments WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	model = &CommentsModel{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&model.Id, &model.ParentId, &model.PostId, &model.Message, &model.Created, &model.State, &model.ArticleIdFk)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}
