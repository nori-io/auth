package postgres

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/errors"
	"gorm.io/gorm"

	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/pkg/transactor"

	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository struct {
	Tx transactor.Transactor
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	m := newModel(e)

	if err := r.Tx.GetDB(ctx).Create(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()

	return nil
}

func (r *UserRepository) Update(ctx context.Context, e *entity.User) error {
	m := newModel(e)
	if err := r.Tx.GetDB(ctx).Save(m).Error; err != nil {
		return errors.NewInternal(err)
	}

	*e = *m.convert()

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Tx.GetDB(ctx).Delete(&model{ID: id}).Error; err != nil {
		return errors.NewInternal(err)
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	out := &model{}
	err := r.Tx.GetDB(ctx).Where("id=?", id).First(out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	out := &model{}

	err := r.Tx.GetDB(ctx).Where("email=?", email).First(&out).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	out := &model{}

	//@todo find by phone number and country code
	err := r.Tx.GetDB(ctx).Where("CONCAT(phone_number, phone_country_code)=?", phone).First(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return out.convert(), nil
}

func (r *UserRepository) FindByFilter(ctx context.Context, filter repository.UserFilter) ([]entity.User, error) {
	var (
		models   []model
		entities []entity.User
	)
	q := r.Tx.GetDB(ctx).Offset(filter.Offset).Limit(filter.Limit)
	if filter.EmailPattern != nil {
		q = q.Where("email LIKE ?", filter.EmailPattern)
	}
	if filter.PhonePattern != nil {
		q = q.Where("CONCAT(phone_number, phone_country_code) LIKE ?", filter.PhonePattern)
	}

	if filter.UserStatus != nil {
		q = q.Where("status = ?", filter.UserStatus.Value())
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

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.Tx.GetDB(ctx).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
