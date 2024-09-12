package validators

import (
	"github.com/gin-gonic/gin"
)

type Validator interface {
	Validate(c *gin.Context) bool
}

const NamePattern string = `^[A-Za-z]{1,50}$`
const DatePattern string = `^\d{4}-\d{2}-\d{2}$`

func IsFormDataValid(c *gin.Context, model interface{}) bool {
	if err := c.ShouldBind(model); err != nil {
		return false
	}
	return true
}

func IsNonEmptyString(s string) bool {
	return len(s) > 0
}

func IsEmpty(s string) bool {
	return len(s) == 0 || s == ""
}

func IsInRange(s string, min, max int) bool {
	return len(s) >= min && len(s) <= max
}
