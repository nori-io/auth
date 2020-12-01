package entity

import "time"

type User struct {
	Id        uint64
	Email     string
	Password  string
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserStatus uint8

const (
	Active UserStatus = iota
	Blocked
	Locked
)

func (u UserStatus) Value() uint8 {
	return uint8(u)
}
func New(status uint8) UserStatus {
	return UserStatus(status)
}
