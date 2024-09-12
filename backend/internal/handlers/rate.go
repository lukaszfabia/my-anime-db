package handlers

import (
	"api/internal/app"
	ratecontroller "api/internal/controller/rate_controller"
	"api/pkg/validators"
	ratevalidator "api/pkg/validators/rate_validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rateValidator validators.Validator = &ratevalidator.RateValidator{}
var rateCtr ratecontroller.RateController = ratecontroller.RateController{}

func AddToList(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if !rateValidator.Validate(c) {
		log.Println("invalid data")
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if err := rateCtr.Save(c, c.Param("id")); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func DeleteFromList(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if err := rateCtr.Delete(c.Param("id")); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}
