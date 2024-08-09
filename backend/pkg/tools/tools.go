package tools

import (
	"api/internal/models"
	"api/pkg/db"
	"log"
	"strconv"
)

func GetOrDefaultNumber(s string, def interface{}) interface{} {
	if s == "" {
		return def
	}

	switch def.(type) {
	case int:
		if res, err := strconv.Atoi(s); err == nil {
			return res
		}
	case float64:
		if res, err := strconv.ParseFloat(s, 64); err == nil {
			return res
		}
	}

	return def
}

type Matchable interface {
	models.AnimeType |
		models.Score |
		models.Pegi |
		models.Status |
		models.CastRole |
		models.GenreOption |
		models.StatusAnime |
		models.FriendRequestStatus
}

/*
Simple matcher parses possible enum value into real enum value

params:

	arr []T - array of possible enum values
	toFind string - string to find in array e.g "tv"
	defVal T - default value if nothing found

returns:

	T - mapped enum value
*/
func Match[T Matchable](arr []T, toFind string, defVal T) T {
	for _, v := range arr {
		if string(v) == toFind {
			return v
		}
	}

	return defVal
}

type parsable interface {
	models.Genre | models.Studio
}

/*
Parse parses map of strings into array of entities

params:

	dicts map[string]string - map of strings to parse e.g map[string]string{"name": "Madhouse"}
	cond string - condition to find entity e.g "name = ?"

returns:

	[]*T - array of entities
*/
func Parse[T parsable](dicts map[string]string, cond string) []*T {
	var result []*T

	for _, value := range dicts {
		var entity T

		if err := db.DB.Where(cond, value).First(&entity).Error; err != nil {
			log.Printf("%s not found in the database", value)
			return nil
		}
		result = append(result, &entity)
	}

	return result
}
