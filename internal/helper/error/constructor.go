package error

import (
	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/helper"
)

type errorHelper struct {
	logger logger.FieldLogger
}

type Params struct {
	Logger logger.FieldLogger
}

func New(params Params) helper.ErrorHelper {
	return &errorHelper{logger: params.Logger}
}
