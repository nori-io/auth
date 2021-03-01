package mfa_type

type MfaType uint8

const (
	OTP MfaType = iota
	Phone
)

func (u MfaType) Value() uint8 {
	return uint8(u)
}

func New(mfaType uint8) MfaType {
	return MfaType(mfaType)
}
