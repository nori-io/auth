package mfa_recovery_codes

import (
	"crypto/rand"
)

func (h mfaRecoveryCodesHelper) Generate(count int) ([]string, error) {
	var codes []string

	for i := 0; i < count; i++ {
		sid := make([]byte, 32)

		if _, err := rand.Read(sid); err != nil {
			return nil, err
		}

		codes = append(codes, string(sid))

	}
	return codes, nil
}
