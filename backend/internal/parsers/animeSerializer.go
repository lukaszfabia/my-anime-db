package parsers

import (
	"api/internal/models"
	"api/pkg/tools"
	"api/pkg/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AnimeToDbFormat(c *gin.Context, anime *models.Anime) (*models.Anime, error) {
	var defEpisodes, defEpisodeLength string = "0", "0"
	var defStartDate, defFinishDate, defDesc, title, alternativeTitles string = "", "", "", "", ""
	var defAnimeType models.AnimeType = models.TV
	var defPegi models.Pegi = models.PG13

	if c.PostForm("title") == "" && anime == nil {
		return nil, errors.New("title is required")
	}

	if anime != nil {
		title = anime.Title
		alternativeTitles = anime.AlternativeTitles
	}

	if anime != nil {
		defEpisodes = strconv.Itoa(anime.Episodes)
		defEpisodeLength = strconv.Itoa(anime.EpisodeLength)
		defStartDate = anime.StartDate
		defFinishDate = anime.FinishDate
		defDesc = anime.Description
		defAnimeType = anime.Type
		defPegi = anime.Pegi
	}

	episodes := tools.GetOrDefaultNumber(c.DefaultPostForm("episodes", defEpisodes), 0).(int)
	episodesLength := tools.GetOrDefaultNumber(c.DefaultPostForm("episodeLength", defEpisodeLength), 0).(int)

	startDate := c.DefaultPostForm("startDate", defStartDate)
	finishDate := c.DefaultPostForm("finishDate", defFinishDate)

	normalizedDesc := strings.ReplaceAll(c.DefaultPostForm("description", defDesc), "'", "")

	var genres []*models.Genre
	var studios []*models.Studio

	if len(c.PostFormMap("genres")) > 0 && anime == nil {
		genres = tools.Parse[models.Genre](c.PostFormMap("genres"), "name = ?")
	} else {
		genres = anime.Genres
	}

	if len(c.PostFormMap("studios")) > 0 && anime == nil {
		studios = tools.Parse[models.Studio](c.PostFormMap("studios"), "name = ?")
	} else {
		studios = anime.Studios
	}

	animeType := tools.Match(models.AllAnimeTypes, c.DefaultPostForm("type", string(defAnimeType)), models.TV)
	pegi := tools.Match(models.AllPegis, c.DefaultPostForm("pegi", string(defPegi)), models.PG13)

	var picUrl string

	if anime != nil {
		utils.RemoveImage(*anime.PicUrl)
	}

	picUrl = utils.SaveImage(c, utils.SaverProps{
		Dir:         utils.AnimesImg,
		Placeholder: utils.DefaultConent,
		KeyToImg:    "picUrl",
		Filename:    c.DefaultPostForm("title", title),
	})

	result := models.Anime{
		Title:             c.DefaultPostForm("title", title),
		AlternativeTitles: c.DefaultPostForm("alternativeTitles", alternativeTitles),
		Type:              animeType,
		Episodes:          episodes,
		Description:       normalizedDesc,
		EpisodeLength:     episodesLength,
		StartDate:         startDate,
		FinishDate:        finishDate,
		PicUrl:            &picUrl,
		Pegi:              pegi,
		Genres:            genres,
		Studios:           studios,
	}

	return &result, nil
}
