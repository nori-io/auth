package entity

import "time"

type ServiceProvider struct {
	ID        uint64
	Name      string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
