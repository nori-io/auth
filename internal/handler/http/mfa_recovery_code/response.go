package mfa_recovery_code

type MfaRecoveryCodesResponse struct {
	success bool     `json:"success"`
	message string   `json:"message"`
	codes   []string `json:"codes"`
}
