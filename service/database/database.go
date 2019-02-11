package database

import (
	"database/sql"
	"sync"
)

type Database interface {
	Article() Article
	Comment() Comment
	User() User
}

type Article interface {
	Create(*ArticlesModel) error
	Update(*ArticlesModel) error
	FindById(id int64) (model *ArticlesModel, err error)
	//FindByEmail(email string) (model *ArticlesModel, err error)
}

type Comment interface {
	Create(*CommentsModel) error
	Update(*CommentsModel) error
	FindById(id int64) (model *CommentsModel, err error)
	//FindByEmail(email string) (model *ArticlesModel, err error)
}

type User interface {
	Create(*UserModel) error
	Update(*UserModel) error
	FindById(id string) (model *UserModel, err error)
	FindByEmail(email string) (model *UserModel, err error)
}

type database struct {
	db      *sql.DB
	article *article
	comment *comment
	user    *user
}

var instance *database
var once sync.Once

// Create Database using singltone pattern
func DB(db *sql.DB) Database {
	once.Do(func() {
		instance = &database{
			db: db,
			article: &article{
				db: db,
			},
			comment: &comment{
				db: db,
			},
			user: &user{
				db: db,
			},
		}
	})
	return instance
}

func (db *database) Article() Article {
	return db.article
}

func (db *database) Comment() Comment {
	return db.comment
}

func (db *database) User() User {
	return db.user
}
