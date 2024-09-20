package voiceactorvalidator

import (
	"api/internal/models"
	"api/pkg/validators"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type VoiceActorValidator struct {
}

func (av *VoiceActorValidator) Validate(c *gin.Context) bool {
	name, lastname, birthdate := c.PostForm("name"), c.PostForm("lastname"), c.PostForm("birthdate")
	nameReg := regexp.MustCompile(validators.NamePattern)
	_, err := time.Parse(time.DateOnly, birthdate)

	if !validators.IsFormDataValid(c, &models.VoiceActor{}) {
		return false
	}

	if c.Request.Method == "POST" {
		return nameReg.MatchString(strings.TrimSpace(name)) && nameReg.MatchString(strings.TrimSpace(lastname)) && err == nil
	}

	if c.Request.Method == "PUT" {
		return (name != "" && nameReg.MatchString(name) || name == "") &&
			(lastname != "" && nameReg.MatchString(lastname) || lastname == "") &&
			(birthdate == "" || err == nil)
	}

	return false
}
