package users_action

type Action uint8

const (
	ChangePassword Action = iota
	DisableMfa
	EnableMfa
	MfaRecoveryCodeApply
	MfaRecoveryCodesGenerate
	LogIn
	LogInMfa
	LogOut
	SignUp
	RestorePassword
)

func (u Action) Value() uint8 {
	return uint8(u)
}

func New(action uint8) Action {
	return Action(action)
}
