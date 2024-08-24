package accountvalidator

import (
	"api/internal/models"
	"api/pkg/validators"
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin"
)

func isSecurePassword(password string, empty bool) bool {
	const midReg string = `^[A-Za-z\d]{8,}$`           // one big letter & one decimal & len = 8
	const strongReg string = `^[A-Za-z\d@$!%#?&]{8,}$` // the same above + special char

	goodPasswordValidator := regexp.MustCompile(midReg)
	strongPasswordValidator := regexp.MustCompile(strongReg)

	var basicCond bool = goodPasswordValidator.MatchString(password) || strongPasswordValidator.MatchString(password)

	if empty {
		return basicCond || password == ""
	}

	return basicCond
}

func isEmailValid(email string, empty bool) bool {
	_, err := mail.ParseAddress(email)
	if empty {
		return err == nil || email == ""
	}
	return err == nil
}

func isUsernameValid(username string, empty bool) bool {
	const nameReg string = `^[A-Za-z\d]{3,50}$`

	validUsername := regexp.MustCompile(nameReg)

	if empty {
		return validUsername.MatchString(username) || username == ""
	}

	return validUsername.MatchString(username)
}

type AccountValidator struct {
}

func (av *AccountValidator) Validate(c *gin.Context) bool {
	username, email, password := c.PostForm("username"), c.PostForm("email"), c.PostForm("password")
	canBeEmpty := c.Request.Method == "PUT"
	basicCond := isSecurePassword(password, canBeEmpty) &&
		isEmailValid(email, canBeEmpty) &&
		isUsernameValid(username, canBeEmpty)

	if c.Request.Method == "POST" {
		return basicCond && validators.IsFormDataValid(c, &models.SignupForm{})
	}

	if c.Request.Method == "PUT" {
		bio, website := c.PostForm("bio"), c.PostForm("website")
		maxLenForBio, maxLenForWebsite := 100, 50

		lenCond := len(bio) <= maxLenForBio && len(website) <= maxLenForWebsite

		return basicCond && lenCond && validators.IsFormDataValid(c, &models.UpdateAccountForm{})
	}

	return false
}
