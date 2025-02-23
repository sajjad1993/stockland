package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
