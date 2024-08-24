package handlers

import (
	"api/internal/models"
	"api/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAnimes(c *gin.Context) {
	var animes models.Anime
	order := db.ToOrder(db.DB, "title", "type", "status")
	if err := db.RetrieveAll(&models.Anime{}, &animes, order,
		db.Association{Model: "AlternativeTitles"},
		db.Association{Model: "AnimeStat"},
		db.Association{Model: "Genres"},
		db.Association{Model: "Studio"},
	); err != nil {
		return
	}

	c.JSON(http.StatusOK, animes)
}

func RetrieveAnime(c *gin.Context) {
	id := c.Param("id")

	var anime models.Anime

	if err := db.Retrieve(&models.Anime{}, &anime, id,
		db.Association{Model: "Role"},
		db.Association{Model: "Genres"},
		db.Association{Model: "Studio"},
		db.Association{Model: "Prequel"},
		db.Association{Model: "Sequel"},
		db.Association{Model: "AlternativeTitles"},
		db.Association{Model: "AnimeStat"},
	); err != nil {
		return
	}

	c.JSON(http.StatusOK, anime)
}

func CreateAnime(c *gin.Context) {
	//TODO: implement
}

func EditAnime(c *gin.Context) {
	// TODO: implement
}

func DeleteAnime(c *gin.Context) {
	//TODO: implement
}
