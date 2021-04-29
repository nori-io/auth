package reset_password

type RestorePasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SetPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
