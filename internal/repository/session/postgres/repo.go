package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type SessionRepository struct {
	Tx transactor.Transactor
}

func (r *SessionRepository) Create(ctx context.Context, e *entity.Session) error {
	modelSession := NewModel(e)

	lastRecord := new(model)

	if err := r.Tx.GetDB(ctx).Create(modelSession).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *SessionRepository) Update(ctx context.Context, e *entity.Session) error {
	model := NewModel(e)
	err := r.Tx.GetDB(ctx).Save(model).Error

	return err
}

func (r *SessionRepository) FindBySessionKey(ctx context.Context, sessionKey string) (*entity.Session, error) {
	var (
		out = &model{}
		e   error
	)
	e = r.Tx.GetDB(ctx).Where("session_key=?", sessionKey).First(out).Error

	return out.Convert(), e
}
