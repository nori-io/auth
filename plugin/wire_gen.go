// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nori-io/common/v4/pkg/domain/registry"
	pg "github.com/nori-io/interfaces/database/orm/gorm"
	http2 "github.com/nori-io/interfaces/nori/http"
	"github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/handler/http"
	"github.com/nori-plugins/authentication/internal/handler/http/authentication"
	mfa_recovery_code4 "github.com/nori-plugins/authentication/internal/handler/http/mfa_recovery_code"
	mfa_secret3 "github.com/nori-plugins/authentication/internal/handler/http/mfa_secret"
	mfa_recovery_code2 "github.com/nori-plugins/authentication/internal/helper/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/authentication_log"
	"github.com/nori-plugins/authentication/internal/repository/mfa_recovery_code"
	"github.com/nori-plugins/authentication/internal/repository/mfa_secret"
	"github.com/nori-plugins/authentication/internal/repository/user"
	"github.com/nori-plugins/authentication/internal/service/auth"
	mfa_recovery_code3 "github.com/nori-plugins/authentication/internal/service/mfa_recovery_code"
	mfa_secret2 "github.com/nori-plugins/authentication/internal/service/mfa_secret"
)

// Injectors from wire.go:

func Initialize(registry2 registry.Registry, config2 config.Config) (*http.Handler, error) {
	httpHttp, err := http2.GetHttp(registry2)
	if err != nil {
		return nil, err
	}
	db, err := pg.GetGorm(registry2)
	if err != nil {
		return nil, err
	}
	userRepository := user.New(db)
	authenticationLogRepository := authentication_log.New(db)
	sessionSession, err := session.GetSession(registry2)
	if err != nil {
		return nil, err
	}
	params := auth.Params{
		UserRepository:              userRepository,
		AuthenticationLogRepository: authenticationLogRepository,
		Session:                     sessionSession,
	}
	authenticationService := auth.New(params)
	mfaRecoveryCodeRepository := mfa_recovery_code.New(db)
	mfa_recovery_codeParams := mfa_recovery_code2.Params{
		Config: config2,
	}
	mfaRecoveryCodesHelper := mfa_recovery_code2.New(mfa_recovery_codeParams)
	params2 := mfa_recovery_code3.Params{
		MfaRecoveryCodeRepository: mfaRecoveryCodeRepository,
		MfaRecoveryCodeHelper:     mfaRecoveryCodesHelper,
		Config:                    config2,
	}
	mfaRecoveryCodeService := mfa_recovery_code3.New(params2)
	authenticationHandler := authentication.New(authenticationService)
	mfaRecoveryCodeHandler := mfa_recovery_code4.New(mfaRecoveryCodeService)
	mfaSecretRepository := mfa_secret.New(db)
	mfa_secretParams := mfa_secret2.Params{
		MfaSecretRepository: mfaSecretRepository,
		UserRepository:      userRepository,
		Config:              config2,
	}
	mfaSecretService := mfa_secret2.New(mfa_secretParams)
	mfaSecretHandler := mfa_secret3.New(mfaSecretService)
	handler := &http.Handler{
		R:                      httpHttp,
		AuthenticationService:  authenticationService,
		MfaRecoveryCodeService: mfaRecoveryCodeService,
		Config:                 config2,
		AuthenticationHandler:  authenticationHandler,
		MfaRecoveryCodeHandler: mfaRecoveryCodeHandler,
		MfaSecretHandler:       mfaSecretHandler,
	}
	return handler, nil
}

// wire.go:

var set = wire.NewSet(pg.GetGorm, http2.GetHttp)
