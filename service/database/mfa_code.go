package database

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

type code struct {
	db  *sql.DB
	Log *log.Logger
}

const (
	chars    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsLen = len(chars)
	mask     = 1<<6 - 1
)

var rng = rand.NewSource(time.Now().UnixNano())

func (c *code) Create(userId uint64, modelMfaCode *MfaCodeModel) error {

	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	for index := 0; index < 10; index++ {
		_, execErr := tx.Exec("INSERT INTO user_mfa_phone (id, user_id, code) VALUES(?,?)",
			userId, RandStr(5)+""+RandStr(5))
		if execErr != nil {
			_ = tx.Rollback()
			return execErr
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

func RandStr(ln int) string {
	buf := make([]byte, ln)
	for idx, cache, remain := ln-1, rng.Int63(), 10; idx >= 0; {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf[idx] = chars[int(cache&mask)%charsLen]
		cache >>= 6
		remain--
		idx--
	}
	return *(*string)(unsafe.Pointer(&buf))
}
