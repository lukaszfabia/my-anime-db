package reviewcontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"errors"
)

type ReviewController struct{}

func (rc *ReviewController) Save(user models.User, animeId, review string) error {

	var collection models.UserAnime

	if err := db.DB.
		Model(models.UserAnime{}).
		Where("anime_id = ? AND user_id = ?", animeId, user.ID).
		First(&collection).Error; err != nil || collection.Status != models.Completed {
		return errors.New("you have to watch this anime first")
	}

	collection.Review = review

	if db.DB.Save(&collection).Error != nil {
		return errors.New("could not save review")
	}

	return nil
}
