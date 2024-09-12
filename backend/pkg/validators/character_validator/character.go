package charactervalidator

import (
	"api/internal/models"
	"api/pkg/validators"
	"regexp"

	"github.com/gin-gonic/gin"
)

type CharacterValidator struct {
}

func (cv *CharacterValidator) Validate(c *gin.Context) bool {
	name, lastname, info := c.PostForm("name"), c.PostForm("lastname"), c.PostForm("information")
	nameReg := regexp.MustCompile(validators.NamePattern)

	if !validators.IsFormDataValid(c, &models.Character{}) {
		return false
	}

	if c.Request.Method == "POST" {
		return nameReg.MatchString(name) &&
			nameReg.MatchString(lastname) &&
			(len(info) <= 10000 && validators.IsNonEmptyString(info))
	}

	if c.Request.Method == "PUT" {
		return (name != "" && nameReg.MatchString(name) || name == "") &&
			(lastname != "" && nameReg.MatchString(lastname) || lastname == "") &&
			(info == "" || (len(info) <= 10000 && validators.IsNonEmptyString(info)))
	}

	return false
}
