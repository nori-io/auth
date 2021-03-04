package config

import "github.com/nori-io/common/v4/pkg/domain/config"

type Config struct {
	UrlPrefix                config.String
	MfaRecoveryCodePattern   config.String
	MfaRecoveryCodeSymbols   config.String
	MfaRecoveryCodeMaxLength config.Int
	MfaRecoveryCodeCount     config.Int
	Issuer                   config.String
}
