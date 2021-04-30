package session

import (
	"context"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"
	"github.com/nori-plugins/authentication/internal/domain/service"

	"github.com/nori-plugins/authentication/pkg/enum/session_status"
)

func (srv SessionService) Create(ctx context.Context, data service.SessionCreateData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	session, err := srv.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil && err != errors2.SessionNotFound {
		return err
	}
	if session != nil && session.Status == session_status.Active {
		return errors2.ActiveSessionAlreadyExists
	}

	if err := srv.sessionRepository.Create(ctx, &entity.Session{
		UserID:     data.UserID,
		SessionKey: []byte(data.SessionKey),
		Status:     data.Status,
		OpenedAt:   data.OpenedAt,
	}); err != nil {
		return err
	}

	return nil
}

func (srv SessionService) Update(ctx context.Context, data service.SessionUpdateData) error {
	if err := data.Validate(); err != nil {
		return err
	}

	session, err := srv.GetBySessionKey(ctx, service.GetBySessionKeyData{SessionKey: data.SessionKey})
	if err != nil {
		return err
	}

	if session == nil {
		return errors2.SessionNotFound
	}

	if err := srv.sessionRepository.Update(ctx, &entity.Session{
		UserID:     data.UserID,
		SessionKey: []byte(data.SessionKey),
		Status:     data.Status,
		ClosedAt:   data.ClosedAt,
		UpdatedAt:  data.UpdatedAt,
	}); err != nil {
		return err
	}
	return nil
}

func (srv SessionService) GetBySessionKey(ctx context.Context, data service.GetBySessionKeyData) (*entity.Session, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	session, err := srv.sessionRepository.FindBySessionKey(ctx, data.SessionKey)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors2.SessionNotFound
	}
	return session, nil
}
