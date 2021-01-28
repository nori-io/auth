package authentication

import (
	"context"
	"time"
)

type Social interface {
	GetAccounts(ctx context.Context, userID uint64) []SocialAccount
	GetAccountsByFilter(ctx context.Context)
	GetAccessToken(ctx context.Context, refreshToken string, serviceProviderID uint64) (accessToken string)
	GetServiceProviders(ctx context.Context) []ServiceProvider
}

type SocialFilter struct {
	ID                  uint64
	ExternalID          string
	FirstName           string
	LastName            string
	FullName            string
	Email               string
	ServiceProviderName string
	Offset              int
	Limit               int
}

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

type ServiceProvider struct {
	ID        uint64
	Name      string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
