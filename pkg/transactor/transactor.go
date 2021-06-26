package transactor

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/errors"

	"gorm.io/gorm"
)

const (
	keyTx = "tx"
)

type Transactor interface {
	GetDB(ctx context.Context) *gorm.DB
	Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error)
}

func (t *TxManager) GetDB(ctx context.Context) *gorm.DB {
	transaction, ok := ctx.Value(keyTx).(*gorm.DB)

	if ok {
		return transaction
	}

	return t.db
}

func (t *TxManager) Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error) {
	var dbTx *gorm.DB

	dbTx, ok := ctx.Value(keyTx).(*gorm.DB)
	if !ok || dbTx == nil {
		dbTx = t.db.Begin()
		ctx = context.WithValue(ctx, keyTx, dbTx)
	}

	defer func() {
		if e := recover(); e != nil {
			err = errors.New("error_recover", err.Error(), errors.ErrInternal)
		}
		if err != nil {
			if e := dbTx.Rollback().Error; e != nil {
				t.log.Error("%s", e)
				return
			}
			return
		}
		if e := dbTx.Commit().Error; e != nil {
			t.log.Error("%s", e)
			return
		}
	}()

	return txFunc(ctx)
}
