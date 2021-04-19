package transactor

import (
	"github.com/google/wire"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

var TransactorSet = wire.NewSet(
	wire.Struct(new(transactor.Params), "Db", "Logger"),
	transactor.New)
