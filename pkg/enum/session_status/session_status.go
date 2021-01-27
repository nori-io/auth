package session_status

type SessionStatus uint8

const (
	Active SessionStatus = iota
	Inactive
)

func (u SessionStatus) Value() uint8 {
	return uint8(u)
}

func New(status uint8) SessionStatus {
	return SessionStatus(status)
}
