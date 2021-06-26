package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/pkg/errors"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type ResetPasswordRepository struct {
	Tx transactor.Transactor
}

func (r *ResetPasswordRepository) Create(ctx context.Context, e *entity.ResetPassword) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()

	return nil
}

func (r *ResetPasswordRepository) Delete(ctx context.Context, userID uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{UserID: userID}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}

func (r *ResetPasswordRepository) FindByToken(ctx context.Context, token string) (*entity.ResetPassword, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("token=?", token).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}
