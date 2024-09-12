package handlers

import (
	"api/internal/app"
	categorycontroller "api/internal/controller/category_controller"
	"api/pkg/validators"
	categoryvalidator "api/pkg/validators/category_validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	r := app.Gin{Ctx: c}
	var v validators.Validator = &categoryvalidator.CategoryValidator{}

	if !v.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	categories := c.QueryArray("category")

	catCtr := &categorycontroller.CategoryController{
		Categories: categories,
	}

	var data any
	var err error

	if categories == nil {
		data, err = catCtr.GetAll()
	} else {
		data, err = catCtr.Get()
	}

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, err)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, data)
}
