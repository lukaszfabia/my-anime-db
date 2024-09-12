package categorycontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"errors"
)

type CategoryController struct {
	Categories []string
}

var Accepted map[string]any = map[string]any{
	"anime_type":   models.AllAnimeTypes,
	"anime_status": models.AllAnimeStatuses,
	"pegi":         models.AllPegis,
	"cast_role":    models.AllCastRoles,
	"score":        models.AllScores,
	"watch_status": models.AllWatchStatuses,
	"genre":        []models.Genre{},
	"studio":       []models.Studio{},
	"character":    []models.Character{},
	"voice_actor":  []models.VoiceActor{},
	"anime":        []models.Anime{},
}

var selectFields = map[string]string{
	"genre":       "id, name",
	"studio":      "id, name",
	"voice_actor": "id, name, last_name, pic_url",
	"character":   "id, name, last_name, pic_url",
	"anime":       "id, title",
}

var orderFields = map[string]string{
	"genre":       "name, id",
	"studio":      "name, id",
	"voice_actor": "last_name, name, id",
	"character":   "last_name, name, id",
	"anime":       "title, id",
}

func (cc *CategoryController) GetAll() (map[string]any, error) {
	if cc.Categories != nil {
		return cc.Get()
	}
	res := make(map[string]any)

	var err error

	for k := range Accepted {
		if fields, ok := selectFields[k]; ok {
			order := orderFields[k]
			res[k], err = find(k, fields, order)

			if err != nil {
				return nil, err
			}

		}
	}

	return res, nil
}

func (cc *CategoryController) Get() (map[string]any, error) {
	res := make(map[string]any)
	var err error

	var col []string = cc.Categories

	for _, category := range col {
		if fields, ok := selectFields[category]; ok {
			order := orderFields[category]
			res[category], err = find(category, fields, order)

			if err != nil {
				return nil, err
			}

		} else {
			res[category] = Accepted[category]
		}
	}

	return res, nil
}

func find(category, fields, order string) (any, error) {
	switch category {
	case "genre":
		var genres []models.Genre
		if err := db.DB.Model(&models.Genre{}).Select(fields).Order(order).Find(&genres).Error; err != nil {
			return nil, errors.New("can not find model")
		}
		return genres, nil
	case "studio":
		var studios []models.Studio
		if err := db.DB.Model(&models.Studio{}).Select(fields).Order(order).Find(&studios).Error; err != nil {
			return nil, errors.New("can not find model")
		}
		return studios, nil
	case "voice_actor":
		var voiceActors []models.VoiceActor
		if err := db.DB.Model(&models.VoiceActor{}).Select(fields).Order(order).Find(&voiceActors).Error; err != nil {
			return nil, errors.New("can not find model")
		}
		return voiceActors, nil
	case "character":
		var characters []models.Character
		if err := db.DB.Model(&models.Character{}).Select(fields).Order(order).Find(&characters).Error; err != nil {
			return nil, errors.New("can not find model")
		}
		return characters, nil

	case "anime":
		var animes []models.Anime
		if err := db.DB.Model(&models.Anime{}).Select(fields).Order(order).Find(&animes).Error; err != nil {
			return nil, errors.New("can not find model")
		}
		return animes, nil

	default:
		return nil, errors.New("problem with db")
	}

}
