package handlers

import (
	"api/internal/models"
	"api/internal/parsers"
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
		Where("title <> ''").
		Take(&animes)

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
	id := c.Param("id")

	var anime models.Anime

	if id == "" {
		anime, e := parsers.AnimeToDbFormat(c, nil)
		if err := db.DB.Create(&anime).Error; err != nil || e != nil {
			msgErr := "Title is required or anime with given title already exists"
			c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
			return
		}
	} else {
		if err := db.DB.Preload("Genres").Preload("Studios").Preload("Roles").First(&anime, id).Error; err != nil {

			msgErr := "No anime found"
			c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
			return
		}

		updatedAnime, _ := parsers.AnimeToDbFormat(c, &anime)
		anime.Title = updatedAnime.Title
		anime.AlternativeTitles = updatedAnime.AlternativeTitles
		anime.Type = updatedAnime.Type
		anime.Episodes = updatedAnime.Episodes
		anime.Description = updatedAnime.Description
		anime.EpisodeLength = updatedAnime.EpisodeLength
		anime.StartDate = updatedAnime.StartDate
		anime.FinishDate = updatedAnime.FinishDate
		anime.Pegi = updatedAnime.Pegi
		anime.Genres = updatedAnime.Genres
		anime.Studios = updatedAnime.Studios

		if err := db.DB.Save(&anime).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	msg := "Anime saved"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}

func RemoveAnime(c *gin.Context) {
	id := c.Param("id")

	err := db.Delete(&models.Anime{}, id, "Genres", "Studios", "Roles")

	if err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Anime removed"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}
