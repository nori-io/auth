package authentication_log

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type AuthenticationLogService struct {
	authenticationLogRepository repository.AuthenticationLogRepository
	transactor                  transactor.Transactor
}

type Params struct {
	AuthenticationLogRepository repository.AuthenticationLogRepository
	Transactor                  transactor.Transactor
}

func New(params Params) service.AuthenticationLogService {
	return &AuthenticationLogService{
		authenticationLogRepository: params.AuthenticationLogRepository,
		transactor:                  params.Transactor,
	}
}
