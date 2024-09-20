package reviewcontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"errors"
)

type ReviewController struct{}

func (rc *ReviewController) Save(user models.User, animeId, review string) (*models.Review, error) {

	var r models.Review

	if err := db.DB.
		Model(models.Review{}).
		Where("anime_id = ? AND user_id = ?", animeId, user.ID).
		Preload("UserAnime").
		FirstOrCreate(&r).Error; err != nil || r.UserAnime.Status != models.Completed {
		return nil, errors.New("you have to watch this anime first")
	}

	tx := db.DB.Begin()

	var stat models.UserStat

	if err := db.DB.First(&stat, "user_id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("could not find user stats")
	}

	if err := stat.AfterPostReview(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	r.Content = review

	if db.DB.Save(&r).Error != nil {
		return nil, errors.New("could not save review")
	}

	return &r, nil
}
