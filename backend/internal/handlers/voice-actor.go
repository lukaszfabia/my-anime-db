package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	voiceactorcontroller "api/internal/controller/voice_actor_controller"
	"api/internal/models"
	"api/pkg/validators"
	voiceactorvalidator "api/pkg/validators/voice_actor_validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var actorController controller.Controller[models.VoiceActor] = &voiceactorcontroller.VoiceActorController{}
var av validators.Validator = &voiceactorvalidator.VoiceActorValidator{}

func GetVoiceActor(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")

	voiceActor, err := actorController.Get(id)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, voiceActor)
}

func GetAllVoiceActors(c *gin.Context) {
	r := app.Gin{Ctx: c}

	voiceActors, err := actorController.GetAll()

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, voiceActors)
}

func CreateVoiceActor(c *gin.Context) {
	var r app.Gin = app.Gin{Ctx: c}
	if !av.Validate(c) {
		log.Println("Validation failed")
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if newActor, err := actorController.Create(c); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusCreated, app.Ok, newActor)
	}
}

func DeleteVoiceActor(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if err := actorController.Delete(id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func EditVoiceActor(c *gin.Context) {
	id := c.Param("id")
	r := app.Gin{Ctx: c}

	if !av.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if editedActor, err := actorController.Update(c, id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusOK, app.Ok, editedActor)
	}

}
