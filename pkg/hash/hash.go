// Package hash that provide helper methods to deal with crypto
package hash

import "golang.org/x/crypto/bcrypt"

func Bcrypt(value string) (string, error) {
	cost := 12

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func BcryptVerify(value, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(value), []byte(hash))
}
