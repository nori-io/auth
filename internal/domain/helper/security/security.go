package security

type SecurityHelper interface {
	GenerateHash(salt, password string) (string, error)
	ComparePassword(salt, hash, password string) (bool, error)
}
