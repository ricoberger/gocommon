package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

// HashString returns a hashed string and an error
func HashString(password string) (string, error) {
	key, err := HashBytes([]byte(password))
	return string(key), err
}

// HashBytes returns a hashed byte array and an error
func HashBytes(password []byte) ([]byte, error) {
	key, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// MatchString returns true if the hash matches the password
func MatchString(hash, password string) bool {
	return MatchBytes([]byte(hash), []byte(password))
}

// MatchBytes returns true if the hash matches the password
func MatchBytes(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err == nil {
		return true
	}

	return false
}
