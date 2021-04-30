package users_action

type Action uint8

const (
	PasswordChanged Action = iota
	PasswordReseted
	MfaDisabled
	MfaOtpEnabled
	MfaPhoneEnabled
	MfaRecoveryCodeApply
	MfaRecoveryCodesGenerate
	LogIn
	LogInMfa
	LogOut
	SignUp
	UserStatusChanged
	UserDeleted
)

func (u Action) Value() uint8 {
	return uint8(u)
}

func New(action uint8) Action {
	return Action(action)
}
