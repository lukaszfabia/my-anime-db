package animecontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/tools"
	"api/pkg/utils"
	"api/pkg/validators"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnimeController struct{}

func (ac *AnimeController) GetAll() ([]*models.Anime, error) {
	var animes []*models.Anime

	if err := db.DB.
		Model(&models.Anime{}).
		Order("title DESC").
		Preload("AnimeStat").
		Preload("AlternativeTitles").
		Preload("Genres").
		Preload("Studio").
		Find(&animes).Error; err != nil {
		return nil, err
	}

	return animes, nil
}

func (ac *AnimeController) Get(id string, props ...any) (*models.Anime, error) {
	var anime models.Anime
	var order string

	if len(props) > 0 {
		order = fmt.Sprintf("CASE WHEN user_id = %s THEN 0 ELSE 1 END, updated_at DESC", props[0])
	} else {
		log.Println("No user ID provided, sorting by updated_at DESC")
		order = "updated_at DESC"
	}

	if err := db.DB.
		Model(&models.Anime{}).
		Preload("AnimeStat").
		Preload("AlternativeTitles").
		Preload("Genres").
		Preload("Studio").
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Character").Preload("VoiceActor")
		}).
		Preload("Prequel").
		Preload("Sequel").
		Preload("Reviews", func(db *gorm.DB) *gorm.DB {
			return db.Order(order).Preload("User")
		}).
		First(&anime, id).Error; err != nil {
		return nil, err
	}

	return &anime, nil
}

func (ac *AnimeController) Create(c *gin.Context) (*models.Anime, error) {
	var newAnime models.Anime

	animeType := tools.Match(models.AllAnimeTypes, c.PostForm("animeType"), models.TV)
	pegi := tools.Match(models.AllPegis, c.PostForm("pegi"), models.PG13)
	status := tools.Match(models.AllAnimeStatuses, c.PostForm("status"), models.Unknown)

	episodes, _ := strconv.Atoi(c.PostForm("episodes"))
	episodeLength, _ := strconv.Atoi(c.PostForm("episodeLength"))

	startDate, _ := time.Parse(time.DateOnly, c.PostForm("startDate"))
	finishDate, err := time.Parse(time.DateOnly, c.PostForm("finishDate"))

	if err != nil {
		finishDate = startDate
	}

	picUrl := utils.SaveImage(c, utils.AnimesImg, "pic")

	studioId := c.PostForm("studio")
	var studio models.Studio
	if err := db.DB.First(&studio, studioId).Error; err != nil {
		log.Println("studio not found")
		return nil, err
	}

	var roles []*models.Role

	if c.PostFormArray("roles") != nil {
		if err := json.Unmarshal([]byte(c.PostForm("roles")), &roles); err != nil {
			return nil, err
		}
	}

	newAnime = models.Anime{
		Title:             c.PostForm("title"),
		Description:       c.PostForm("description"),
		AlternativeTitles: []*models.OtherTitles{},
		Type:              animeType,
		Pegi:              pegi,
		Status:            status,
		Episodes:          episodes,
		EpisodeLength:     episodeLength,
		StartDate:         &startDate,
		FinishDate:        &finishDate,
		PicUrl:            picUrl,
		Genres:            []*models.Genre{},
		StudioID:          studio.ID,
		Studio:            &studio,
		AnimeStat:         &models.AnimeStat{},
		Roles:             []*models.Role{},
	}

	if err := db.DB.Create(&newAnime).Error; err != nil {
		if picUrl != nil {
			utils.RemoveImage(picUrl)
		}
		return nil, err
	}

	if newAnime, err := Save(c, &newAnime, false); err != nil {
		return nil, err
	} else {
		return newAnime, nil
	}
}

func Save(c *gin.Context, newAnime *models.Anime, update bool) (*models.Anime, error) {
	genres := c.PostFormArray("genres")
	prequelId := c.PostForm("prequel")
	sequelId := c.PostForm("sequel")

	alternativeTitles := c.PostFormArray("altTitles")

	var roles []*models.Role

	for _, altTitle := range alternativeTitles {
		if update {
			if db.DB.Unscoped().Delete(&models.OtherTitles{}, "anime_id = ?", newAnime.ID).Error != nil {
				return nil, errors.New("cannot delete alternative titles")
			}
		}

		otherTitle := models.OtherTitles{
			AlternativeTitle: altTitle,
			AnimeID:          newAnime.ID,
		}
		if err := db.DB.Create(&otherTitle).Error; err != nil {
			return nil, err
		}
		newAnime.AlternativeTitles = append(newAnime.AlternativeTitles, &otherTitle)
	}

	if c.PostFormArray("roles") != nil {
		if err := json.Unmarshal([]byte(c.PostForm("roles")), &roles); err != nil {
			return nil, err
		}
	}

	for _, role := range roles {
		role.AnimeID = newAnime.ID
		var va models.VoiceActor
		var c models.Character

		if err := db.DB.First(&va, role.ActorID).Error; err != nil {
			return nil, err
		}

		if err := db.DB.First(&c, role.CharacterID).Error; err != nil {
			return nil, err
		}

		role.ActorID = va.ID
		role.CharacterID = c.ID
		role.VoiceActor = va
		role.Character = c
		role.Anime = *newAnime

		newAnime.Roles = append(newAnime.Roles, role)
	}

	// TODO : to check but imo it's alright
	var prequel, sequel models.Anime
	if !validators.IsEmpty(prequelId) {
		if err := db.DB.Preload("Sequel").Preload("Prequel").First(&prequel, prequelId).Error; err == nil {
			newAnime.Prequel = &prequel
			newAnime.PrequelID = &prequel.ID
		}
	}

	if !validators.IsEmpty(sequelId) {
		if err := db.DB.First(&sequel, sequelId).Error; err == nil {
			newAnime.Sequel = &sequel
			newAnime.SequelID = &sequel.ID
		}
	}

	for _, genreId := range genres {
		var genre models.Genre
		if err := db.DB.First(&genre, genreId).Error; err != nil {
			log.Println("genre not found")
			return nil, err
		}
		newAnime.Genres = append(newAnime.Genres, &genre)
	}

	if err := db.DB.Save(&newAnime).Error; err != nil {
		return nil, err
	}

	return newAnime, nil
}

func (ac *AnimeController) Update(c *gin.Context, id string) (*models.Anime, error) {
	var animeToUpdate models.Anime

	if err := db.DB.First(&animeToUpdate, id).Error; err != nil {
		return nil, err
	}

	animeToUpdate.Title = controller.GetOrDefault(c.PostForm("title"), animeToUpdate.Title).(string)
	animeToUpdate.Description = controller.GetOrDefault(c.PostForm("description"), animeToUpdate.Description).(string)

	animeToUpdate.Type = tools.Match(models.AllAnimeTypes, c.PostForm("animeType"), animeToUpdate.Type)
	animeToUpdate.Pegi = tools.Match(models.AllPegis, c.PostForm("pegi"), animeToUpdate.Pegi)
	animeToUpdate.Status = tools.Match(models.AllAnimeStatuses, c.PostForm("status"), animeToUpdate.Status)

	animeToUpdate.Episodes = controller.GetOrDefault(c.PostForm("episodes"), animeToUpdate.Episodes).(int)
	animeToUpdate.EpisodeLength = controller.GetOrDefault(c.PostForm("episodeLength"), animeToUpdate.EpisodeLength).(int)

	// try to parse dates
	if sDate, err := time.Parse(time.DateOnly, c.PostForm("startDate")); err == nil {
		animeToUpdate.StartDate = &sDate
	}

	if fDate, err := time.Parse(time.DateOnly, c.PostForm("finishDate")); err == nil {
		animeToUpdate.FinishDate = &fDate
	}

	if newPicUrl := utils.SaveImage(c, utils.AnimesImg, "pic"); newPicUrl != nil {
		if animeToUpdate.PicUrl != nil {
			utils.RemoveImage(animeToUpdate.PicUrl)
		}
		animeToUpdate.PicUrl = newPicUrl
	}

	if studioId := c.PostForm("studio"); studioId != "" {
		var newStudio models.Studio
		if db.DB.First(&newStudio, studioId).Error == nil {
			animeToUpdate.StudioID = newStudio.ID
			animeToUpdate.Studio = &newStudio
		}
	}

	if err := db.DB.Save(&animeToUpdate).Error; err != nil {
		return nil, err
	}

	if animeToUpdate, err := Save(c, &animeToUpdate, true); err != nil {
		return nil, err
	} else {
		return animeToUpdate, nil
	}

}

func (ac *AnimeController) Delete(id string) error {
	if err := db.DB.Unscoped().Delete(&models.AnimeStat{}, "anime_id = ?", id).Error; err != nil {
		return errors.New("cannot delete anime stat")
	}

	if err := db.DB.Unscoped().Delete(&models.Role{}, "anime_id = ?", id).Error; err != nil {
		return errors.New("cannot delete roles")
	}

	if err := db.DB.Unscoped().Delete(&models.OtherTitles{}, "anime_id = ?", id).Error; err != nil {
		return errors.New("cannot delete alternative titles")
	}

	return db.Delete(&models.Anime{}, id,
		db.Association{Model: "Genres"},
		db.Association{Model: "Prequel"},
		db.Association{Model: "Sequel"},
	)
}
