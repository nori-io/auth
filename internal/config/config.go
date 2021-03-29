package config

import "github.com/nori-io/common/v4/pkg/domain/config"

type Config struct {
	UrlPrefix              config.String
	EmailVerification      config.Bool
	EmailActivationCodeTTL config.UInt64
	MfaRecoveryCodePattern config.String
	MfaRecoveryCodeSymbols config.String
	MfaRecoveryCodeLength  config.Int
	MfaRecoveryCodeCount   config.Int
	Issuer                 config.String
	PasswordBcryptCost     config.Int
	CookiesPath            config.String
	CookiesDomain          config.String
	CookiesExpires         config.UInt64
	CookiesRawExpires      config.String
	CookiesMaxAge          config.UInt64
	CookiesSecure          config.Bool
	CookiesHttpOnly        config.Bool
	CookiesSameSite        config.UInt64
	CookiesRaw             config.String
	CookiesUnparsed        config.SliceString
}
