package service

import (
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg"
)

type Service struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
}

type Params struct {
	AuthenticationService  service.AuthenticationService
	MfaRecoveryCodeService service.MfaRecoveryCodeService
}

func New(params Params) pkg.Authentication {
	return &Service{
		AuthenticationService:  params.AuthenticationService,
		MfaRecoveryCodeService: params.MfaRecoveryCodeService,
	}
}
