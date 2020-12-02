package user_status

type UserStatus uint8

const (
	Active UserStatus = iota
	Blocked
	Locked
)

func (u UserStatus) Value() uint8 {
	return uint8(u)
}
func New(status uint8) UserStatus {
	return UserStatus(status)
}
