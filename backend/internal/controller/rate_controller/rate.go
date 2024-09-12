package ratecontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/tools"
	"api/pkg/validators"
	"log"

	"github.com/gin-gonic/gin"
)

type RateController struct{}

func (rc *RateController) Save(c *gin.Context, animeId string) error {
	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		return err
	}

	// checking if anime exists
	var anime models.Anime
	if err := db.DB.Preload("AnimeStat").First(&anime, animeId).Error; err != nil {
		return err
	}

	var elem models.UserAnime
	// update or create new user anime
	if err := db.DB.First(&elem, "user_id = ? AND anime_id = ?", user.ID, anime.ID).Error; err != nil {
		// new elem
		elem = models.UserAnime{
			UserID:  user.ID,
			AnimeID: anime.ID,
		}
	}

	if validators.IsNonEmptyString(c.PostForm("score")) {
		elem.Score = tools.Match(models.AllScores, c.PostForm("score"), models.Good)
	}
	if validators.IsNonEmptyString(c.PostForm("watchStatus")) {
		elem.Status = tools.Match(models.AllWatchStatuses, c.PostForm("watchStatus"), models.Watching)
	}

	if validators.IsNonEmptyString(c.PostForm("isFav")) {
		elem.IsFav = c.PostForm("isFav") == "true"
	}

	if err := db.DB.Save(&elem).Error; err != nil {
		return err
	}

	tx := db.DB.Begin()
	var stat models.AnimeStat
	if err := tx.First(&stat, "anime_id = ?", anime.ID).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	var a models.UserAnime
	var new bool

	if err := tx.First(&a, "user_id = ? AND anime_id = ?", user.ID, anime.ID).Error; err != nil {
		return err
	}

	new = a.CreatedAt.Equal(a.UpdatedAt) // if it's new

	if err := stat.AfterAddAnime(tx, new); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (rc *RateController) Delete(id string) error {
	return db.DB.Delete(&models.UserAnime{}, id).Error
}
