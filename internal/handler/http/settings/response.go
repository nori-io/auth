package settings

type DisableMfaResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MfaStatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type PasswordChangeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
