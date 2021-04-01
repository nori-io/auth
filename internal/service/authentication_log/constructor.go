package authentication_log

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationLogService struct {
	authenticationLogRepository repository.AuthenticationLogRepository
	db                          *gorm.DB
}

type Params struct {
	authenticationLogRepository repository.AuthenticationLogRepository
	DB                          *gorm.DB
}

func New(params Params) service.AuthenticationLogService {
	return &AuthenticationLogService{
		authenticationLogRepository: params.authenticationLogRepository,
		db:                          params.DB,
	}
}
