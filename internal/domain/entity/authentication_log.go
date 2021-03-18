package entity

import (
	"time"

	"github.com/nori-plugins/authentication/pkg/enum/users_action"
)

type AuthenticationLog struct {
	ID        uint64
	UserID    uint64
	Action    users_action.Action
	Meta      string
	CreatedAt time.Time
}
