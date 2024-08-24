package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	charactercontroller "api/internal/controller/character_controller"
	"api/internal/models"
	"api/pkg/validators"
	charactervalidator "api/pkg/validators/character_validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var characterController controller.Controller[models.Character] = &charactercontroller.CharacterController{}
var cv validators.Validator = &charactervalidator.CharacterValidator{}

func GetAllCharacters(c *gin.Context) {
	r := app.Gin{Ctx: c}
	characters, err := characterController.GetAll()

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, characters)
}

func GetCharacter(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")
	character, err := characterController.Get(id)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, character)
}

func CreateCharacter(c *gin.Context) {
	r := app.Gin{Ctx: c}
	if !cv.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
	}

	if newCharacter, err := characterController.Create(c); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusCreated, app.Ok, newCharacter)
	}
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if err := characterController.Delete(id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func EditCharacter(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if !cv.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if edtiedCharacter, err := characterController.Update(c, id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusOK, app.Ok, edtiedCharacter)
	}
}
