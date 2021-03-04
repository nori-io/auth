package repository

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/repository"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/mfa_secret"
	"github.com/nori-plugins/authentication/internal/repository/user"
)

var RepositorySet = wire.NewSet(
	user.New,
	mfa_recovery_code.New,
	mfa_secret.New,
	repository.New,
)
