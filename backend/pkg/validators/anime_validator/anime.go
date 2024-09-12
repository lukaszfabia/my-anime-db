package animevalidator

import (
	"api/internal/models"
	"api/pkg/tools"
	"api/pkg/validators"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AnimeValidator struct{}

func (av *AnimeValidator) Validate(c *gin.Context) bool {
	// get values
	animeType, status, pegi := c.PostForm("animeType"), c.PostForm("status"), c.PostForm("pegi")
	isEnumValid :=
		(tools.CheckEnum(models.AllAnimeTypes, animeType) &&
			tools.CheckEnum(models.AllAnimeStatuses, status) &&
			tools.CheckEnum(models.AllPegis, pegi))

	s, f := c.PostForm("startDate"), c.PostForm("finishDate")
	episodes, duration := c.PostForm("episodes"), c.PostForm("episodeLength")

	// checking
	if !validators.IsEmpty(episodes) {
		_, err := strconv.Atoi(episodes)
		return err == nil
	}

	if !validators.IsEmpty(duration) {
		_, err := strconv.Atoi(duration)
		return err == nil
	}

	if validators.IsNonEmptyString(c.PostForm("sequel")) && validators.IsNonEmptyString(c.PostForm("prequel")) {
		if c.PostForm("sequel") == c.PostForm("prequel") {
			return false
		}
	}

	start, errS := time.Parse(time.DateOnly, s)
	finish, errF := time.Parse(time.DateOnly, f)
	// only if both dates are not empty
	if !validators.IsEmpty(s) && !validators.IsEmpty(f) {
		parsableDate := errF == nil && errS == nil
		return parsableDate && !start.After(finish)
	}

	// nonsense when finsh date exists and start date doesnt
	if !validators.IsEmpty(f) && validators.IsEmpty(s) {
		return false
	}

	if !isEnumValid {
		return false
	}

	if c.Request.Method == "POST" {
		return validators.IsFormDataValid(c, &models.AnimeForm{})
	}

	return c.Request.Method == "PUT"
}
