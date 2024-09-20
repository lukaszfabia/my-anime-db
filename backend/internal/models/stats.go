package models

import (
	"errors"
	"log"
	"sort"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AnimeStat struct {
	gorm.Model
	AnimeID          uint    `gorm:"primaryKey" json:"animeId"`
	GlobalScore      float64 `gorm:"default:0.0" json:"score"`
	Popularity       uint    `gorm:"default:0" json:"popularity"`
	MostPopularGrade Score   `gorm:"type:text;default" json:"mostPopularGrade"`
}

type UserStat struct {
	gorm.Model
	UserID        uint           `gorm:"primaryKey" json:"userId"`
	WatchedHours  float64        `gorm:"default:0.0" json:"watchedHours"` // only completed animes
	PostedReviews *int64         `gorm:"default:0" json:"postedReviews"`
	FavGenres     pq.StringArray `gorm:"type:text[]" json:"favGenres"`
}

func (us *UserStat) UpdateUserStats(tx *gorm.DB, props ...any) error {
	var ratedAnime []*UserAnime

	if err := tx.Model(&UserAnime{}).Preload("Anime", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Genres")
	}).Where("user_id = ?", us.UserID).Find(&ratedAnime).Error; err != nil || len(ratedAnime) == 0 {
		return errors.New("could not find anime stats")
	}

	var sum float64
	for _, e := range ratedAnime {
		if e.Status == Completed {
			sum += float64(e.Anime.EpisodeLength) * float64(e.Anime.Episodes)
		}
	}
	us.WatchedHours = sum / 60

	var genres map[string]int = make(map[string]int)
	var maxCard int = 0

	for _, e := range ratedAnime {
		if e.Status == Completed {
			for _, g := range e.Anime.Genres {
				genres[g.Name]++
				if genres[g.Name] > maxCard {
					maxCard = genres[g.Name]
				}
			}
		}
	}

	var favGenres pq.StringArray

	for k, v := range genres {
		if v == maxCard {
			favGenres = append(favGenres, k)
		}
	}

	us.FavGenres = favGenres

	if err := tx.Save(&us).Error; err != nil {
		log.Println(err)
		return errors.New("could not save stats")
	}

	return nil
}

func (us *UserStat) AfterPostReview(tx *gorm.DB) error {
	if err := tx.Model(&Review{}).Where("user_id = ?", us.UserID).Count(us.PostedReviews).Error; err != nil {
		return errors.New("could not count reviews")
	}

	if err := tx.Save(&us).Error; err != nil {
		return errors.New("could not save stats")
	}

	return nil
}

func (a *AnimeStat) UpdateAnimeStats(tx *gorm.DB) error {
	var ratedAnime []*UserAnime

	if err := tx.Model(&UserAnime{}).Where("anime_id = ?", a.AnimeID).Find(&ratedAnime).Error; err != nil || len(ratedAnime) == 0 {
		return errors.New("could not find anime stats")
	}

	var cs *ComputeScores = NewComputeScores(ratedAnime)

	a.GlobalScore, a.MostPopularGrade = cs.AvgScore(), cs.GetPopularGrade()
	if err := tx.Save(a).Error; err != nil {
		return errors.New("could not save stats")
	}

	if err := cs.SetPopularity(tx); err != nil {
		return errors.New("could not set popularity")
	}

	return nil
}

type ComputeScores struct {
	ratedAnime []*UserAnime
	scoreMap   map[Score]int
	scores     map[Score]int
}

func NewComputeScores(ratedAnime []*UserAnime) *ComputeScores {
	var res map[Score]int = make(map[Score]int, len(AllScores))
	for weight, score := range AllScores {
		res[score] = weight + 1
	}

	scores := make(map[Score]int, len(ratedAnime))

	return &ComputeScores{ratedAnime: ratedAnime, scoreMap: res, scores: scores}
}

func (cs *ComputeScores) AvgScore() float64 {

	var sum int

	for _, stat := range cs.ratedAnime {
		sum += cs.scoreMap[stat.Score]
		cs.scores[stat.Score]++
	}

	avg := float64(sum) / float64(len(cs.ratedAnime))
	return avg
}

func (cs *ComputeScores) GetPopularGrade() Score {

	popularGrade := Bad
	for score, count := range cs.scores {
		if cs.scores[popularGrade] < count {
			popularGrade = score
		}
	}
	return popularGrade
}

func (cs *ComputeScores) SetPopularity(tx *gorm.DB) error {
	var allAnimesFormLists []*UserAnime
	var allAnimesStats []*AnimeStat

	if err := tx.Find(&allAnimesFormLists).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Find(&allAnimesStats).Error; err != nil {
		log.Println(err)
		return err
	}

	// only for sort animes by cardinality
	idAndCardinality := make(map[uint]int)

	for _, e := range allAnimesFormLists {
		idAndCardinality[e.AnimeID]++
	}

	var ids []uint
	for k := range idAndCardinality {
		ids = append(ids, k)
	}

	sort.SliceStable(ids, func(i, j int) bool {
		return idAndCardinality[ids[i]] > idAndCardinality[ids[j]]
	})

	log.Println(ids)

	for _, anime := range allAnimesStats {
		for index, id := range ids {
			if id == anime.AnimeID {
				anime.Popularity = uint(index + 1) // offset cuz indexing starts from 0
				if err := tx.Save(&anime).Error; err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}
	return nil
}
