package transactor

import (
	"context"

	"github.com/nori-plugins/authentication/pkg/errors"

	"github.com/jinzhu/gorm"
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
	var tx *gorm.DB

	tx, ok := ctx.Value(keyTx).(*gorm.DB)
	if !ok || tx == nil {
		tx = t.db.Begin()
		ctx = context.WithValue(ctx, keyTx, tx)
	}

	defer func() {
		if e := recover(); e != nil {
			err = errors.New("error_recover", err.Error(), errors.ErrInternal)
		}
		if err != nil {
			if e := tx.Rollback().Error; e != nil {
				t.log.Error("%s", e)
				return
			}
			return
		}
		if e := tx.Commit().Error; e != nil {
			t.log.Error("%s", e)
			return
		}
	}()

	return txFunc(ctx)
}
