package authentication

import "context"

type Social interface {
	GetAccounts(ctx context.Context, userID uint64) []SocialAccount
	GetAccessToken(ctx context.Context, refreshToken string, serviceProviderID uint64) (accessToken string)
	GetServiceProviders(ctx context.Context) []ServiceProvider
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
}

type ServiceProvider struct {
	ID   uint64
	Name string
	Logo string
}
