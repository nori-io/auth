package postgres

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/nori-plugins/authentication/internal/domain/entity"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/pkg/transactor"
)

type SocialProviderRepository struct {
	Tx transactor.Transactor
}

func (r SocialProviderRepository) Create(ctx context.Context, e *entity.SocialProvider) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()
	return nil
}

func (r SocialProviderRepository) Update(ctx context.Context, e *entity.SocialProvider) error {
	m := newModel(e)
	if err := r.Tx.GetDB(ctx).Save(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()
	return nil
}

func (r SocialProviderRepository) Delete(ctx context.Context, ID uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{ID: ID}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}

func (r SocialProviderRepository) FindByID(ctx context.Context, ID uint64) (*entity.SocialProvider, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("id=?", ID).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r SocialProviderRepository) FindByName(ctx context.Context, name string) (*entity.SocialProvider, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("name=?", name).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r SocialProviderRepository) FindByFilter(ctx context.Context, filter repository.SocialProviderFilter) ([]entity.SocialProvider, error) {
	var (
		models   []model
		entities []entity.SocialProvider
	)
	q := r.Tx.GetDB(ctx).Offset(filter.Offset).Limit(filter.Limit)
	if filter.Status != nil {
		q = q.Where("status LIKE ?", filter.Status)
	}

	err := q.Find(&models).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}
	for _, v := range models {
		entities = append(entities, *v.convert())
	}
	return entities, nil
}
