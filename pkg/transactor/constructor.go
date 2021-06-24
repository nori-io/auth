package transactor

import (
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"gorm.io/gorm"
)

type TxManager struct {
	db  *gorm.DB
	log logger.FieldLogger
}

type Params struct {
	Db     *gorm.DB
	Logger logger.FieldLogger
}

func New(params Params) Transactor {
	return &TxManager{
		db:  params.Db,
		log: params.Logger,
	}
}
