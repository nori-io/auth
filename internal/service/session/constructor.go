package session

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SessionService struct {
	sessionRepository repository.SessionRepository
}

type Params struct {
	SessionRepository repository.SessionRepository
}

func New(params Params) service.SessionService {
	return &SessionService{sessionRepository: params.SessionRepository}
}
