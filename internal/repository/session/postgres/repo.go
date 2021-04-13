package postgres

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionRepository struct {
	Tx transactor.Transactor
}

func (r *SessionRepository) Create(ctx context.Context, e *entity.Session) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}
	*e = *m.convert()

	return nil
}

func (r *SessionRepository) Update(ctx context.Context, e *entity.Session) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Save(m).Error; err != nil {
		return errors.NewInternal(err)
	}
	*e = *m.convert()

	return nil
}

func (r *SessionRepository) FindBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("session_key=?", sessionKey).First(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}
