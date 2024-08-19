package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// for list
type ShortAnime struct {
	Title         string           `json:"title"`
	Type          models.AnimeType `json:"type"`
	Episodes      int              `json:"episodes"`
	Description   string           `json:"description"`
	EpisodeLength int              `json:"episodeLength"`
	Pegi          models.Pegi      `json:"pegi"`
	PicUrl        string           `json:"picUrl"`
	GlobalScore   float64          `json:"score"`
	Popularity    uint             `json:"popularity"`
}

func GetAllAnimes(c *gin.Context) {
	var animes []ShortAnime

	res := db.DB.Model(&models.Anime{}).
		Find(&animes)

	if res.Error != nil {
		msgErr := "There is no any anime"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))

		return
	}

	c.JSON(http.StatusOK, animes)
}

func RetrieveAnime(c *gin.Context) {
	id := c.Param("id")

	var anime models.Anime

	res := db.DB.
		Preload("Genres").Preload("Studios").Preload("Roles").
		Where("id = ? AND title <> ''", id).
		First(&anime)

	if res.Error != nil {
		msgErr := "No anime found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))

		return
	}

	c.JSON(http.StatusOK, anime)
}

func SaveOrUpdateAnime(c *gin.Context) {

}

func RemoveAnime(c *gin.Context) {
}
