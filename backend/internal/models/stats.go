package models

import (
	"errors"
	"log"
	"sort"

	"gorm.io/gorm"
)

type AnimeStat struct {
	gorm.Model
	AnimeID          uint    `gorm:"primaryKey" json:"animeId"`
	GlobalScore      float64 `gorm:"default:0.0" json:"score"`
	Popularity       uint    `gorm:"default:0" json:"popularity"`
	MostPopularGrade Score   `gorm:"type:text;default" json:"mostPopularGrade"`
}

func (a *AnimeStat) AfterAddAnime(tx *gorm.DB, new bool) error {
	var ratedAnime []*UserAnime

	if err := tx.Model(&UserAnime{}).Where("anime_id = ? AND status = ?", a.AnimeID, Completed).Find(&ratedAnime).Error; err != nil || len(ratedAnime) == 0 {
		return errors.New("could not find anime stats")
	}

	var cs *ComputeScores = NewComputeScores(ratedAnime)

	if new {
		cs.SetPopularity(tx)
	}

	a.GlobalScore, a.MostPopularGrade = cs.AvgScore(), cs.GetPopularGrade()

	if err := tx.Save(a).Error; err != nil {
		return errors.New("could not save stats")
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
	log.Println(avg)
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

	if err := tx.Model(&UserAnime{}).Where("status = ?", Completed).Find(&allAnimesFormLists).Error; err != nil {
		return err
	}

	if err := tx.Find(&allAnimesStats).Error; err != nil {
		return err
	}

	// only for sort animes by cardinality
	idAndCardinality := make(map[uint]int)

	for _, e := range allAnimesFormLists {
		// e.AnimeID
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
				anime.Popularity = uint(index + 1) // offset cuz indexing starting from 0
				if err := tx.Save(&anime).Error; err != nil {
					return errors.New("could not save anime popularity")
				}
			}
		}
	}

	return nil
}
