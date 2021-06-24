package helper

type MfaRecoveryCodesHelper interface {
	Generate() ([]string, error)
}
