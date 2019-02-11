package service

import (

	"github.com/sirupsen/logrus"
	"github.com/nori-io/auth/service/database"
)

type Service interface {
}

type service struct {
	db  database.Database
	log *logrus.Logger
}

func NewService(
	log *logrus.Logger,
	db database.Database,
) Service {
	return &service{

	db: db,

		log: log,
	}
}
