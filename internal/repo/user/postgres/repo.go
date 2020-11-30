package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nori-io/authentication/internal/domain/entity"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) (*entity.User, error) {

	model, _ := NewModel(e)

	lastRecord := new(User)

	err := r.Db.Create(model).Scan(&lastRecord).Error
	l := lastRecord.Convert()
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (r *UserRepository) Get(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		out = &User{}
		e   error
	)
	if r.Db == nil {
		return nil, errors.New("database error")
	}
	e = r.Db.Where("ID=?", id).First(out).Error

	return out.Convert(), e
}

func (r *UserRepository) GetAll(ctx context.Context, offset uint64, limit uint64) ([]entity.User, error) {
	var (
		out         []User
		outEntities []entity.User
		e           error
	)
	e = r.Db.Find(&out).Error
	for i, v := range out {
		outEntities = append(outEntities, *v.Convert())
		fmt.Println("OUT is", outEntities[i])

	}
	return outEntities, e
}

func (r *UserRepository) Update(ctx context.Context, e *entity.User) error {
	model, _ := NewModel(e)
	err := r.Db.Save(model).Error

	fmt.Println("Put error", err)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {

	if err := r.Db.Delete(&User{Id: id}).Error; err != nil {
		return err
	}
	return nil
}
