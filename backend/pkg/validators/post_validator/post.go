package postvalidator

import (
	"api/internal/models"
	"api/pkg/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostValidator struct{}

func (av *PostValidator) Validate(c *gin.Context) bool {
	title, content := c.PostForm("title"), c.PostForm("content")
	maxTitleLen, maxContentLen := 100, 5000
	basicCond :=
		(len(title) <= maxTitleLen &&
			len(content) <= maxContentLen &&
			validators.IsNonEmptyString(title) &&
			validators.IsNonEmptyString(content))

	if c.Request.Method == "POST" {
		// title and content [1, 100] and [1, 5000] respectively

		return validators.IsFormDataValid(c, &models.PostForm{}) && basicCond
	}

	if c.Request.Method == "PUT" {
		_, err := strconv.ParseBool(c.PostForm("isPublic"))
		parsableBool := err == nil
		return parsableBool && (basicCond || validators.IsEmpty(title) || validators.IsEmpty(content))
	}

	return false
}
