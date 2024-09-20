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

	var elem models.UserAnime = models.UserAnime{
		UserID:  user.ID,
		AnimeID: anime.ID,
	}

	var review models.Review = models.Review{
		UserID:  user.ID,
		AnimeID: anime.ID,
	}

	if err := db.DB.FirstOrCreate(&review, "user_id = ? AND anime_id = ?", user.ID, anime.ID).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := db.DB.FirstOrCreate(&elem, "user_id = ? AND anime_id = ?", user.ID, anime.ID).Error; err != nil {
		log.Println(err)
		return err
	}

	elem.Status = tools.Match(models.AllWatchStatuses, c.PostForm("watchStatus"), models.Watching)

	if validators.IsNonEmptyString(c.PostForm("score")) {
		elem.Score = tools.Match(models.AllScores, c.PostForm("score"), "")
	}

	if elem.Status == models.PlanToWatch {
		elem.Score = ""
	}

	if validators.IsNonEmptyString(c.PostForm("isFav")) {
		elem.IsFav = c.PostForm("isFav") == "true"
	}

	if err := db.DB.Save(&elem).Error; err != nil {
		return err
	}

	tx := db.DB.Begin()
	var stat models.AnimeStat
	var userStat models.UserStat

	if err := tx.First(&userStat, "user_id = ?", user.ID).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := tx.First(&stat, "anime_id = ?", anime.ID).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := userStat.UpdateUserStats(tx); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	if err := stat.UpdateAnimeStats(tx); err != nil {
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
