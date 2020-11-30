package entity

import "time"

// @TODO add json tags?
type User struct {
	Id            uint64
	Email         string
	Password      string
	ProfileTypeId int64
	StatusId      int64
	Kind          string
	Created       time.Time
	Updated       time.Time
}
