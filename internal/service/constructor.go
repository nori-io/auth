package service

import (
	s "github.com/nori-io/interfaces/nori/session"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg"
)

type Service struct {
	session                s.Session
	authenticationService  service.AuthenticationService
	mfaRecoveryCodeService service.MfaRecoveryCodeService
	config                 config
}

type config struct {
}

func New(session s.Session,
	authenticationService service.AuthenticationService,
	mfaRecoveryCodeService service.MfaRecoveryCodeService,
) pkg.Authentication {
	return &Service{
		session:                session,
		authenticationService:  authenticationService,
		mfaRecoveryCodeService: mfaRecoveryCodeService,
		config:                 config{},
	}
}
