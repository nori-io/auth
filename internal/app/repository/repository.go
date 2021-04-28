package repository

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/repository"
	"github.com/nori-plugins/authentication/internal/repository/authentication_log"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/mfa_totp"
	"github.com/nori-plugins/authentication/internal/repository/one_time_token"
	"github.com/nori-plugins/authentication/internal/repository/session"
	"github.com/nori-plugins/authentication/internal/repository/social_account"
	"github.com/nori-plugins/authentication/internal/repository/social_provider"
	"github.com/nori-plugins/authentication/internal/repository/user"
)

var RepositorySet = wire.NewSet(
	authentication_log.New,
	mfa_recovery_code.New,
	mfa_totp.New,
	one_time_token.New,
	social_provider.New,
	session.New,
	social_account.New,
	user.New,
	repository.New,
)
