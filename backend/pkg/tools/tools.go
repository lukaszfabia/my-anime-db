package tools

import (
	"api/internal/models"
	"api/pkg/db"
	"bytes"
	"errors"
	"html/template"
	"log"
	"path"
)

type Matchable interface {
	models.AnimeType |
		models.Score |
		models.Pegi |
		models.WatchStatus |
		models.CastRole |
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

func CheckEnum[T Matchable](arr []T, toFind string) bool {
	for _, v := range arr {
		if string(v) == toFind {
			return true
		}
	}

	return false
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

func ParseHTMLToString(templateName string, data any) (string, error) {
	templatePath := path.Join("../templates/emails", templateName)
	tmpl, err := template.ParseFiles("../templates/base.html", templatePath)

	if err != nil {
		log.Printf("Error parsing template files: %v", err)
		return "", errors.New("failed to parse email template")
	}

	var buf bytes.Buffer

	err = tmpl.ExecuteTemplate(&buf, "base.html", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return "", errors.New("failed to execute email template")
	}

	body := buf.String()
	if body == "" {
		log.Println("Email body is empty")
		return "", errors.New("email body is empty")
	}

	return body, nil
}

func Any[T comparable](arr []T, toFind T) bool {
	for _, e := range arr {
		if e == toFind {
			return true
		}
	}
	return false
}
