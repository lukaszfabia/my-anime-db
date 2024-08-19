package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RetrieveActor(c *gin.Context) {
	id := c.Param("id")
	var actor models.VoiceActor

	res := db.DB.Model(&models.VoiceActor{}).
		Where("voice_actors.id = ?", id).
		Where("voice_actors.name <> ''").
		Joins("JOIN roles ON roles.actor_id = voice_actors.id").
		Joins("JOIN characters ON characters.id = roles.character_id").
		Joins("JOIN animes ON animes.id = roles.anime_id").
		Select("voice_actors.*, roles.*, characters.*, animes.*").
		First(&actor)

	if res.Error != nil {
		msgErr := "No actors"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, actor)
}

func GetAllActors(c *gin.Context) {

}
