package transactor

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/jinzhu/gorm"
)

type Transactor interface {
	GetDB(ctx context.Context) *gorm.DB
	Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error)
}

func (t *TxManager) GetDB(ctx context.Context) *gorm.DB {
	if t.db != nil {
		return t.db
	}
	return ctx.Value("db")
}

func (t *TxManager) Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error) {
	if t.db == nil {
		t.db.Begin()
	}

	NewCtx := context.WithValue(ctx, "db", t)

	defer func() {
		if e := recover(); e != nil {
			err = errors.New("error_recover", err.Error(), errors.ErrInternal)
		}
		if err != nil {
			if e := t.db.Rollback().Error; e != nil {
				t.log.Error("%s", e)
				return
			}
			return
		}
		if e := t.db.Commit().Error; e != nil {
			t.log.Error("%s", e)
			return
		}
	}()

	return txFunc(NewCtx)
}
