package entity

import "time"

type User struct {
	Id        uint64
	Email     string
	Password  string
	Status    int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
