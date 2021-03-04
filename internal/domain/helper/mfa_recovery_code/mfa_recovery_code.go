package mfa_recovery_code

type MfaRecoveryCodesHelper interface {
	Generate(count int) ([]string, error)
}
