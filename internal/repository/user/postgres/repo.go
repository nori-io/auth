package postgres

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/nori-plugins/authentication/internal/domain/entity"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) error {
	model := NewModel(e)

	lastRecord := new(User)

	if err := r.Db.Create(model).Scan(&lastRecord).Error; err != nil {
		return err
	}
	lastRecord.Convert()

	return nil
}

func (r *UserRepository) Get(ctx context.Context, id uint64) (*entity.User, error) {
	var (
		out = &User{}
		e   error
	)
	e = r.Db.Where("id=?", id).First(out).Error

	return out.Convert(), e
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var (
		out = &User{}
		e   error
	)
	e = r.Db.Where("email=?", email).First(out).Error

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
	model := NewModel(e)
	err := r.Db.Save(model).Error

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	if err := r.Db.Delete(&User{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
