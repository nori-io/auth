package reset_password

type ResetPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResetPasswordSetResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
