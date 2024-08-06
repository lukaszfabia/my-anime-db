package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {

	if password == "" {
		return nil, errors.New("password is empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, errors.New("couldnt hash password")
	}

	return hashed, nil
}
