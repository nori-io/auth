package authentication

import "context"

type Social interface {
	GetUserSocialAccounts(ctx context.Context, userID uint64) []SocialAccount
	IsValid(ctx context.Context, tokenAccess string) bool

	RefreshToken(ctx context.Context, tokenRefresh string) (tokenAccess string)
}

type SocialAccount struct {
	name        string
	logo        string
	authApi     string
	callBackApi string
}
