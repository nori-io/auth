package database

import (
	"time"
)

type ArticlesModel struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	State           string `json:"state"`
	MetaDescription string `json:"meta_description"`
	Tags            string `json:"tags"`
}

type CommentsModel struct {
	Id          int64     `json:"id"`
	ParentId    int64     `json:"-"`
	PostId      int64     `json:"-"`
	Message     string    `json:"message"`
	Created     time.Time `json:"created"`
	State       string    `json:"state"`
	ArticleIdFk int64     `json:"article_id_fk"`
}

type UserModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"-"`

	Phone            string `json:"-"`
	PhoneCountryCode string `json:"-"`
	PhoneVerified    bool   `json:"-"`

	Salt     string `json:"-"`
	Password string `json:"-"`

	State string `json:"-"`

	MfaEnabled bool `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
