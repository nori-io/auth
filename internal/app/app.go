package app

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/internal/app/handler"
	"github.com/nori-plugins/authentication/internal/app/helper"
	"github.com/nori-plugins/authentication/internal/app/repository"
	"github.com/nori-plugins/authentication/internal/app/service"
	"github.com/nori-plugins/authentication/internal/app/transactor"
)

var AppSet = wire.NewSet(transactor.TransactorSet, handler.HandlerSet, helper.HelperSet, repository.RepositorySet, service.ServiceSet)
