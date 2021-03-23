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
}
