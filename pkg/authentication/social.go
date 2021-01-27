package authentication

import "context"

type Social interface {
	GetUserSocialAccounts(ctx context.Context, userID uint64) []SocialAccount
	IsValid(ctx context.Context, tokenAccess string) bool
	RefreshToken(ctx context.Context, tokenRefresh string) (tokenAccess string)
}

type SocialAccount struct {
	UserProfile   UserProfile
	SocialNetwork SocialNetwork
}

type UserProfile struct {
	ID        uint64
	FirstName string
	LastName  string
	FullName  string
	Email     string
	AvatarURL string
	Raw       string
}

type SocialNetwork struct {
	name string
	logo string
}
