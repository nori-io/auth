package database

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const (
	// scrypt is used for strong keys
	// these are the recommended scrypt parameters
	scryptN      = 16384
	scryptR      = 8
	scryptP      = 1
	scryptKeyLen = 32
)

func CreateSalt() ([]byte, error) {
	// Generate a salt
	salt := make([]byte, 65)
	_, err := rand.Read(salt)
	return salt, err
}

func hmacSha256(in, salt []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, salt)
	_, err := mac.Write(in)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

func createKey(in, salt []byte) ([]byte, error) {
	return scrypt.Key(in, salt, scryptN, scryptR, scryptP, scryptKeyLen)
}

func Hash(password, salt []byte) ([]byte, error) {
	//	bytes, _ := CreateSalt()
	peppered, _ := hmacSha256(password, salt)
	cur, _ := createKey(peppered, salt)
	return cur, nil
}

func VerifyPassword(password, salt, hash []byte) (bool, error) {
	h, _ := Hash(password, salt)

	if subtle.ConstantTimeCompare(h, hash) != 1 {

		return false, fmt.Errorf("Invalid username or password")
	}

	return true, nil
}
