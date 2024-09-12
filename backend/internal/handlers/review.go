package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	reviewcontroller "api/internal/controller/review_controller"
	"api/pkg/validators"
	reviewvalidator "api/pkg/validators/review_validator"

	"net/http"

	"github.com/gin-gonic/gin"
)

var revCtr reviewcontroller.ReviewController = reviewcontroller.ReviewController{}
var revValid validators.Validator = &reviewvalidator.ReviewValidator{}

/*
Save a review
*/
func SetReview(c *gin.Context) {
	user, err := controller.GetUserFromCtx(c)
	r := app.Gin{Ctx: c}

	if err != nil || !revValid.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.Failed, nil)
		return
	}

	if revCtr.Save(user, c.Param("id"), c.PostForm("review")) != nil {
		r.NewResponse(http.StatusBadRequest, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}
