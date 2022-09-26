package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetPasswordHash(password, salt string) (string, error) {
	passwordByte := []byte(password + salt)

	newHash256 := sha256.New()
	if _, err := newHash256.Write(passwordByte); err != nil {
		return "", err
	}

	shaPassword := newHash256.Sum(nil)

	return hex.EncodeToString(shaPassword), nil
}
