package session

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
	"github.com/nori-plugins/authentication/pkg/errors"
)

func (srv SessionService) Create(ctx context.Context, data *entity.Session) error {
	session, err := srv.GetBySessionKey(ctx, string(data.SessionKey))
	if err != nil && err != errors2.SessionNotFound {
		return err
	}
	if session != nil {
		return errors2.SessionAlreadyExists
	}

	if err := srv.transactor.Transact(ctx, func(txCtx context.Context) error {
		if err := srv.sessionRepository.Create(txCtx, data); err != nil {
			return errors.NewInternal(err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (srv SessionService) IsSessionExist(ctx context.Context, sessionKey string) (bool, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return false, errors.NewInternal(err)
	}
	if session != nil && session.Status == session_status.Active {
		return true, errors2.SessionAlreadyExists
	}
	return false, nil
}

func (srv SessionService) GetBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors2.SessionNotFound
	}
	return session, nil
}
