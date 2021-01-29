package entity

import "time"

type SocialAccount struct {
	ID              uint64
	ExternalID      string
	FirstName       string
	LastName        string
	FullName        string
	Email           string
	AvatarURL       string
	ServiceProvider ServiceProvider
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
