package mfa_recovery_code

type MfaRecoveryCodesHelper interface {
	Generate() ([]string, error)
}
