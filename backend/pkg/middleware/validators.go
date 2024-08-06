package middleware

import (
	"net/mail"
	"regexp"
)

func IsSecurePassword(password string) bool {
	const midReg string = `^[A-Za-z\d]{8,}$`           // one big letter & one decimal & len = 8
	const strongReg string = `^[A-Za-z\d@$!%#?&]{8,}$` // the same above + special char

	goodPasswordValidator := regexp.MustCompile(midReg)
	strongPasswordValidator := regexp.MustCompile(strongReg)

	return goodPasswordValidator.MatchString(password) || strongPasswordValidator.MatchString(password)
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsUsernameValid(username string) bool {
	const nameReg string = `^[A-Za-z\d]{3,}$`

	validUsername := regexp.MustCompile(nameReg)

	return validUsername.MatchString(username)
}
