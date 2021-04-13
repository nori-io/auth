package session

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
)

func (srv SessionService) Create(ctx context.Context, data *entity.Session) error {
	session, err := srv.GetBySessionKey(ctx, string(data.SessionKey))
	if err != nil && err != errors2.SessionNotFound {
		return err
	}
	if session != nil && session.Status == session_status.Active {
		return errors2.ActiveSessionAlreadyExists
	}

	if err := srv.sessionRepository.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv SessionService) Update(ctx context.Context, data *entity.Session) error {
	if err := srv.sessionRepository.Update(ctx, data); err != nil {
		return err
	}
	return nil
}

func (srv SessionService) IsActiveSessionExist(ctx context.Context, sessionKey string) (bool, error) {
	session, err := srv.sessionRepository.FindBySessionKey(ctx, sessionKey)
	if err != nil {
		return false, err
	}
	if session != nil && session.Status == session_status.Active {
		return true, nil
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
