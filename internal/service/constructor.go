package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg"
)

type Service struct {
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
}

func New(params Params) pkg.Authentication {
	return &Service{
		authenticationService:  params.AuthenticationService,
		mfaRecoveryCodeService: params.MfaRecoveryCodeService,
	}
}
