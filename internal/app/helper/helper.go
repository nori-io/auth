package helper

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/helper"
	cookieHelper "github.com/nori-plugins/authentication/internal/helper/cookie"
	errorHelper "github.com/nori-plugins/authentication/internal/helper/error"
	gothProviderHelper "github.com/nori-plugins/authentication/internal/helper/goth_provider"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_code"
	securityHelper "github.com/nori-plugins/authentication/internal/helper/security"
)

var HelperSet = wire.NewSet(
	wire.Struct(new(mfaRecoveryCodeHelper.Params), "Config"),
	mfaRecoveryCodeHelper.New,
	wire.Struct(new(cookieHelper.Params), "Config"),
	cookieHelper.New,
	wire.Struct(new(errorHelper.Params), "Logger"),
	errorHelper.New,
	wire.Struct(new(securityHelper.Params), "Config"),
	securityHelper.New,
	helper.New,
	gothProviderHelper.New,
)
