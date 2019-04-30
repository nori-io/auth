package database

import (
	"context"
	"database/sql"
	"math/rand"
	"time"
	"unsafe"

	"github.com/nori-io/nori-common/interfaces"
)

type mfaCode struct {
	db  *sql.DB
	log interfaces.Logger
}

const (
	chars    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsLen = len(chars)
	mask     = 1<<6 - 1
)

var rng = rand.NewSource(time.Now().UnixNano())

func (c *mfaCode) Create(modelMfaCode *MfaCodeModel) ([]string, error) {

	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	var recoveryCodes []string
	if err != nil {
		return nil, err
	}
	_, execErr := tx.Exec("DELETE FROM user_mfa_code WHERE user_шd=?", modelMfaCode.UserId)

	for index := 0; index < 10; index++ {
		generatedCode := RandStr(5) + "-" + RandStr(5)
		_, execErr = tx.Exec("INSERT INTO user_mfa_code (user_id, code) VALUES(?,?)",
			modelMfaCode.UserId, generatedCode)
		if len(generatedCode) != 0 {
			recoveryCodes = append(recoveryCodes, generatedCode)
		}
		if execErr != nil {
			_ = tx.Rollback()
			return nil, execErr
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return recoveryCodes, nil

}

func (c *mfaCode) Delete(code string) error {

	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	_, execErr := tx.Exec("DELETE FROM user_mfa_code WHERE code=?", code)

	if execErr != nil {
		_ = tx.Rollback()
		return execErr
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
