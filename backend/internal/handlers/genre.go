package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	genrecontroller "api/internal/controller/genre_controller"
	"api/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var genreController controller.Controller[models.Genre] = &genrecontroller.GenreController{}

// var av validators.Validator = &voiceactorvalidator.VoiceActorValidator{}

func GetAllGenres(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if allGenres, err := genreController.GetAll(); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, err)
	} else {
		r.NewResponse(http.StatusOK, app.Ok, allGenres)
	}
}

func GetGenre(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")

	if genre, err := genreController.Get(id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, err)
	} else {
		r.NewResponse(http.StatusOK, app.Ok, genre)
	}
}

func CreateGenre(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if genre, err := genreController.Create(c); err != nil {
		log.Println(err)
		r.NewResponse(http.StatusInternalServerError, app.Failed, err)
	} else {
		r.NewResponse(http.StatusOK, app.Ok, genre)
	}
}

func EditGenre(c *gin.Context) {
	r := app.Gin{Ctx: c}

	r.NewResponse(http.StatusOK, app.Failed, nil)
}

func DeleteGenre(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")

	if err := genreController.Delete(id); err != nil {
		log.Println(err)
		r.NewResponse(http.StatusInternalServerError, app.Failed, err)
	} else {
		r.NewResponse(http.StatusOK, app.Ok, nil)
	}
}
