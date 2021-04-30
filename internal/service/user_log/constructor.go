package user_log

import (
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type UserLogService struct {
	userLogRepository repository.UserLogRepository
	transactor        transactor.Transactor
}

type Params struct {
	UserLogRepository repository.UserLogRepository
	Transactor        transactor.Transactor
}

func New(params Params) service.UserLogService {
	return &UserLogService{
		userLogRepository: params.UserLogRepository,
		transactor:        params.Transactor,
	}
}
