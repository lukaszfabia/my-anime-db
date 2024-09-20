package ratevalidator

import (
	"api/internal/models"
	"api/pkg/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RateValidator struct{}

func (rv *RateValidator) Validate(c *gin.Context) bool {
	checkEnum := tools.CheckEnum(models.AllWatchStatuses, c.PostForm("watchStatus"))
	_, err := strconv.ParseBool(c.PostForm("isFav"))

	return checkEnum && err == nil
}
