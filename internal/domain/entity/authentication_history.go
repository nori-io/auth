package entity

import "time"

type AuthenticationHistory struct {
	ID      uint64
	UserID  uint64
	Signin  time.Time
	Meta    string
	Signout time.Time
}
