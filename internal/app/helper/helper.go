package helper

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/helper"
	cookieHelper "github.com/nori-plugins/authentication/internal/helper/cookie"
	errorHelper "github.com/nori-plugins/authentication/internal/helper/error"
	gothProviderHelper "github.com/nori-plugins/authentication/internal/helper/goth_provider"
	mfaRecoveryCodeHelper "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_code"
	mfaTotpHelper "github.com/nori-plugins/authentication/internal/helper/mfa_totp"

	securityHelper "github.com/nori-plugins/authentication/internal/helper/security"
)

var HelperSet = wire.NewSet(
	wire.Struct(new(cookieHelper.Params),
		"Config"),
	cookieHelper.New,
	wire.Struct(new(errorHelper.Params),
		"Logger"),
	errorHelper.New,
	gothProviderHelper.New,
	wire.Struct(new(mfaRecoveryCodeHelper.Params),
		"Config"),
	mfaRecoveryCodeHelper.New,
	wire.Struct(new(mfaTotpHelper.Params),
		"Config"),
	mfaTotpHelper.New,
	wire.Struct(new(securityHelper.Params),
		"Config"),
	securityHelper.New,
	helper.New,
)
