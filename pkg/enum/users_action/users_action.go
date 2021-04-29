package users_action

type Action uint8

const (
	PasswordChanged Action = iota
	PasswordRestored
	MfaDisabled
	MfaEnabled
	MfaRecoveryCodeApply
	MfaRecoveryCodesGenerate
	LogIn
	LogInMfa
	LogOut
	SignUp
)

func (u Action) Value() uint8 {
	return uint8(u)
}

func New(action uint8) Action {
	return Action(action)
}
