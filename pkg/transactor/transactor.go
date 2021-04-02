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
	transaction, ok := ctx.Value("tx").(*gorm.DB)

	if ok {
		return transaction
	}

	return t.db
}

func (t *TxManager) Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error) {
	if t.tx == nil {
		t.tx = t.db.Begin()
	}

	NewCtx := context.WithValue(ctx, "tx", t.tx)

	defer func() {
		if e := recover(); e != nil {
			err = errors.New("error_recover", err.Error(), errors.ErrInternal)
		}
		if err != nil {
			if e := t.tx.Rollback().Error; e != nil {
				t.log.Error("%s", e)
				return
			}
			return
		}
		if e := t.tx.Commit().Error; e != nil {
			t.log.Error("%s", e)
			return
		}
	}()

	return txFunc(NewCtx)
}
