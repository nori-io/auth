package entity

import "time"

type SocialAccount struct {
	ID                uint64
	User_ID           uint64
	ExternalID        string
	ServiceProviderID uint8
	ProviderUserKey   string
	FirstName         string
	LastName          string
	FullName          string
	Email             string
	AvatarURL         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
