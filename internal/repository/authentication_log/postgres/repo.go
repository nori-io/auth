package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/pkg/errors"
)

type AuthenticationLogRepository struct {
	Tx transactor.Transactor
}

func (r *AuthenticationLogRepository) Create(ctx context.Context, e *entity.AuthenticationLog) error {
	modelAuthenticationLog := NewModel(e)

	lastRecord := new(model)

	if err := r.Tx.GetDB(ctx).Create(modelAuthenticationLog).Scan(&lastRecord).Error; err != nil {
		return errors.NewInternal(err)
	}
	lastRecord.Convert()

	return nil
}

func (r *AuthenticationLogRepository) Update(ctx context.Context, e *entity.AuthenticationLog) error {
	model := NewModel(e)
	err := r.Tx.GetDB(ctx).Save(model).Error
	if err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r *AuthenticationLogRepository) FindByUserId(ctx context.Context, userId uint64) (*entity.AuthenticationLog, error) {
	var (
		out = &model{}
		e   error
	)
	e = r.Tx.GetDB(ctx).Where("user_id=?", userId).Last(out).Error

	return out.Convert(), e
}

func (r *AuthenticationLogRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
