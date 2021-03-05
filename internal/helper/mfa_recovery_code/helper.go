package mfa_recovery_code

import (
	"math/rand"
)

func (h mfaRecoveryCodesHelper) Generate() ([]string, error) {
	var codes []string

	for i := 0; i < h.config.MfaRecoveryCodeCount(); i++ {
		code := randomString([]rune(h.config.MfaRecoveryCodeSymbols()), h.config.MfaRecoveryCodeLength())

		codes = append(codes, code)

	}
	return codes, nil
}

func randomString(mfaRecoveryCodeSymbols []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = mfaRecoveryCodeSymbols[rand.Intn(len(mfaRecoveryCodeSymbols))]
	}
	return string(b)
}
