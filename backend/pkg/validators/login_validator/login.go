package loginvalidator

import (
	"api/internal/models"
	"api/pkg/validators"

	"github.com/gin-gonic/gin"
)

type LoginValidator struct{}

func (lv *LoginValidator) Validate(c *gin.Context) bool {
	return validators.IsFormDataValid(c, &models.LoginForm{})
}
