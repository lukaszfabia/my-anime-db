package voiceactorvalidator

import (
	"api/internal/models"
	"api/pkg/validators"
	"regexp"

	"github.com/gin-gonic/gin"
)

type VoiceActorValidator struct {
}

func (av *VoiceActorValidator) Validate(c *gin.Context) bool {
	name, lastname, birthdate := c.PostForm("name"), c.PostForm("lastname"), c.PostForm("birthdate")
	nameReg := regexp.MustCompile(validators.NamePattern)
	dateReg := regexp.MustCompile(validators.DatePattern)

	if !validators.IsFormDataValid(c, &models.VoiceActor{}) {
		return false
	}

	if c.Request.Method == "POST" {
		return nameReg.MatchString(name) && nameReg.MatchString(lastname) && dateReg.MatchString(birthdate)
	}

	if c.Request.Method == "PUT" {
		return (name != "" && nameReg.MatchString(name) || name == "") &&
			(lastname != "" && nameReg.MatchString(lastname) || lastname == "") &&
			(birthdate == "" || dateReg.MatchString(birthdate))
	}

	return false
}
