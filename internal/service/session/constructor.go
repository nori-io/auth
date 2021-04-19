package session

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SessionService struct {
	sessionRepository repository.SessionRepository
	transactor        transactor.Transactor
}

type Params struct {
	SessionRepository repository.SessionRepository
	Transactor        transactor.Transactor
}

func New(params Params) service.SessionService {
	return &SessionService{
		sessionRepository: params.SessionRepository,
		transactor:        params.Transactor,
	}
}
