package settings

type DisableMfaResponse struct {
	Success bool
	Message string
}

type ReceiveMfaResponse struct {
	Success bool
	Message string
	Status  bool
}
