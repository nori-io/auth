package postgres

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/pkg/errors"
)

type UserLogRepository struct {
	Tx transactor.Transactor
}

func (r *UserLogRepository) Create(ctx context.Context, e *entity.UserLog) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()
	return nil
}

func (r *UserLogRepository) Update(ctx context.Context, e *entity.UserLog) error {
	m := newModel(e)
	if err := r.Tx.GetDB(ctx).Save(m).Error; err != nil {
		return errors.NewInternal(err)
	}
	*e = *m.convert()

	return nil
}

func (r *UserLogRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{ID: id}).Error; err != nil {
		errors.NewInternal(err)
	}
	return nil
}

func (r *UserLogRepository) FindByUserID(ctx context.Context, userId uint64) (*entity.UserLog, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("user_id=?", userId).Last(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}
