package security

type SecurityHelper interface {
	GenerateHash(password string) ([]byte, error)
	ComparePassword(hash, password string) error
}
