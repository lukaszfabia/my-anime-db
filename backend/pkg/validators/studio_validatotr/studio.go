package studiovalidatotr

import (
	"api/pkg/validators"

	"github.com/gin-gonic/gin"
)

type StudioValidator struct{}

func (sv *StudioValidator) Validate(c *gin.Context) bool {
	name, date := c.PostForm("name"), c.PostForm("establishedDate")

	return validators.IsNonEmptyString(name) && validators.IsNonEmptyString(date)
}
