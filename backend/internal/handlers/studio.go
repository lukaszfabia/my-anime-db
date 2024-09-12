package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	studiocontroller "api/internal/controller/studio_controller"
	"api/internal/models"
	"api/pkg/validators"
	studiovalidatotr "api/pkg/validators/studio_validatotr"
	"net/http"

	"github.com/gin-gonic/gin"
)

var studioCtr controller.Controller[models.Studio] = &studiocontroller.StudioController{}
var sv validators.Validator = &studiovalidatotr.StudioValidator{}

func GetStudio(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")

	studio, err := studioCtr.Get(id)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, studio)
}

func GetAllStudios(c *gin.Context) {
	r := app.Gin{Ctx: c}

	studios, err := studioCtr.GetAll()

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, studios)
}

func CreateStudio(c *gin.Context) {
	var r app.Gin = app.Gin{Ctx: c}
	if !sv.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if studio, err := studioCtr.Create(c); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusCreated, app.Ok, studio)
	}
}

func DeleteStudio(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if err := studioCtr.Delete(id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func EditStudio(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if studio, err := studioCtr.Update(c, id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusOK, app.Ok, studio)
	}

}
