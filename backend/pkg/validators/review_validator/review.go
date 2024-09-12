package reviewvalidator

import (
	"api/pkg/validators"

	"github.com/gin-gonic/gin"
)

type ReviewValidator struct{}

func (rv *ReviewValidator) Validate(c *gin.Context) bool {
	return validators.IsInRange(c.PostForm("review"), 3, 10000)
}
