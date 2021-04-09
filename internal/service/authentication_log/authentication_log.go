package authentication_log

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

func (srv AuthenticationLogService) Create(ctx context.Context, authenticationLog *entity.AuthenticationLog) error {
	if err := srv.transactor.Transact(ctx, func(tx context.Context) error {
		if err := srv.authenticationLogRepository.Create(tx, authenticationLog); err != nil {
			return errors.NewInternal(err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
