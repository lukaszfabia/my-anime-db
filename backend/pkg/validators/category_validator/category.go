package categoryvalidator

import (
	categorycontroller "api/internal/controller/category_controller"
	"api/pkg/tools"

	"github.com/gin-gonic/gin"
)

type CategoryValidator struct{}

func (cv *CategoryValidator) Validate(c *gin.Context) bool {
	categories := c.QueryArray("category")

	accepted := make([]string, 0, len(categorycontroller.Accepted))
	for key := range categorycontroller.Accepted {
		accepted = append(accepted, key)
	}

	for _, category := range categories {
		if !tools.Any(accepted, category) {
			return false
		}
	}

	return true
}
