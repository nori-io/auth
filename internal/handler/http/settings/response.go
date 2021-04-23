package settings

type DisableMfaResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ReceiveMfaResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
