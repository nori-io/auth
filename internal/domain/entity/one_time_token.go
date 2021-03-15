package entity

import "time"

type OneTimeToken struct {
	ID     uint64
	UserID uint64
	Issuer string
	Token  []byte
	// @todo time.duration? for ttl
	TTL       time.Time
	CreatedAt time.Time
	UsedAt    time.Time
}
