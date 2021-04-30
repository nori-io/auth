package repository

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/mfa_totp"
	"github.com/nori-plugins/authentication/internal/repository/one_time_token"
	"github.com/nori-plugins/authentication/internal/repository/reset_password"
	"github.com/nori-plugins/authentication/internal/repository/session"
	"github.com/nori-plugins/authentication/internal/repository/social_account"
	"github.com/nori-plugins/authentication/internal/repository/social_provider"
	"github.com/nori-plugins/authentication/internal/repository/user"
	"github.com/nori-plugins/authentication/internal/repository/user_log"
)

var RepositorySet = wire.NewSet(
	user_log.New,
	mfa_recovery_code.New,
	mfa_totp.New,
	one_time_token.New,
	reset_password.New,
	session.New,
	social_account.New,
	social_provider.New,
	user.New,
	repository.New,
)
