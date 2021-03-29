package authentication

type SignInResponse struct {
	Success bool
	Message string
	MfaType string
}

type SignInMfaResponse struct {
	Success bool
	Message string
}
