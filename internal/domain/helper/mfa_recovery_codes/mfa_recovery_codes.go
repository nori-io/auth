package mfa_recovery_codes

type MfaRecoveryCodesHelper interface {
	Generate(count int) ([]string, error)
}
