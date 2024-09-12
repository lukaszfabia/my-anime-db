package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	animecontroller "api/internal/controller/anime_controller"
	"api/internal/models"
	"api/pkg/validators"
	animevalidator "api/pkg/validators/anime_validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var animeValidator validators.Validator = &animevalidator.AnimeValidator{}
var animeCtr controller.Controller[models.Anime] = &animecontroller.AnimeController{}

func GetAllAnimes(c *gin.Context) {
	r := app.Gin{Ctx: c}

	animes, err := animeCtr.GetAll()
	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, animes)
}

func GetAnime(c *gin.Context) {
	r := app.Gin{Ctx: c}

	animeId := c.Query("id")

	userId := c.Query("userId")

	var err error

	var anime *models.Anime

	if userId != "" {
		anime, err = animeCtr.Get(animeId, userId)
	} else {
		anime, err = animeCtr.Get(animeId)
	}

	if err != nil {
		log.Println(err)
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, anime)
}

func CreateAnime(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if !animeValidator.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	anime, err := animeCtr.Create(c)
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusCreated, app.Ok, anime)
}

func UpdateAnime(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if !animeValidator.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	id := c.Param("id")

	anime, err := animeCtr.Update(c, id)
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, anime)
}

func DeleteAnime(c *gin.Context) {
	r := app.Gin{Ctx: c}

	id := c.Param("id")

	if err := animeCtr.Delete(id); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}
