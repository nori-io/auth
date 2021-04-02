package transactor

import (
	"github.com/jinzhu/gorm"
	"github.com/nori-io/common/v4/pkg/domain/logger"
)

type TxManager struct {
	db  *gorm.DB
	log logger.FieldLogger
}

type Params struct {
	Db  *gorm.DB
	Log logger.FieldLogger
}

func New(params Params) Transactor {
	return &TxManager{
		db:  params.Db,
		log: params.Log,
	}
}
