package security

func (s securityHelper) GenerateHash(salt, password string) (string, error) {
	panic("implement me")
}

func (s securityHelper) ComparePassword(salt, hash, password string) (bool, error) {
	panic("implement me")
}
