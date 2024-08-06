package models

import (
	"time"

	"gorm.io/gorm"
)

// orm

type User struct {
	gorm.Model
	Username   string       `gorm:"unique;not null" json:"username"`
	Email      string       `gorm:"unique;not null" json:"email"`
	Password   string       `gorm:"not null" json:"password"`
	PicUrl     string       `json:"picUrl"`
	IsVerified bool         `gorm:"default:false" json:"isVerified"`
	IsMod      bool         `gorm:"default:false" json:"isMod"`
	Bio        string       `gorm:"default:'Edit bio'" json:"bio"`
	Website    string       `json:"website"`
	Posts      []*Post      `gorm:"many2many:users_posts;" json:"posts"`
	Friends    []*User      `gorm:"many2many:users_friends;" json:"friends"`
	UserAnimes []*UserAnime `gorm:"foreignKey:UserID" json:"userAnimes"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title"`
	Content  string `gorm:"not null" json:"content"`
	Image    string `json:"image"`
	IsPublic bool   `gorm:"default:true" json:"isPublic"`
}

type Studio struct {
	gorm.Model
	Name    string `gorm:"not null;unique" json:"name"`
	Website string `json:"website"`
}

type Genre struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}

type Anime struct {
	gorm.Model
	Title             string     `gorm:"not null;unique" json:"title"`
	AlternativeTitles string     `json:"alternativeTitles"`
	Type              AnimeType  `gorm:"type:anime_type;default:'tv'" json:"type"`
	Episodes          int        `gorm:"default:0" json:"episodes"`
	Description       string     `gorm:"not null" json:"description"`
	EpisodeLength     int        `gorm:"default:24" json:"episodeLength"`
	StartDate         *time.Time `json:"startDate"`
	FinishDate        *time.Time `json:"finishDate"`
	Pegi              Pegi       `gorm:"type:pegi;default:'PG-13'" json:"pegi"`
	PicUrl            string     `json:"picUrl"`
	GlobalScore       float64    `gorm:"default:0.0" json:"score"`
	Popularity        uint       `gorm:"default:0" json:"popularity"`
	Genres            []*Genre   `gorm:"many2many:anime_genres;" json:"genres"`
	Studios           []*Studio  `gorm:"many2many:anime_studios;" json:"studios"`
}

type VoiceActor struct {
	gorm.Model
	Name      string     `gorm:"not null" json:"name"`
	LastName  string     `gorm:"not null" json:"lastName"`
	PicUrl    string     `json:"picUrl"`
	BirthDate *time.Time `json:"birthDate"`
}

type Character struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Information string `gorm:"default:'update info'" json:"information"`
	PicUrl      string `json:"picUrl"`
}

type UserAnime struct {
	gorm.Model
	UserID  uint   `gorm:"primaryKey" json:"userId"`
	AnimeID uint   `gorm:"primaryKey" json:"animeId"`
	Score   Score  `gorm:"type:score;default:'good'" json:"score"`
	Status  Status `gorm:"type:status;default:'plan-to-watch'" json:"status"`
	Review  string `json:"review"`
}

type Role struct {
	gorm.Model
	ActorID     uint     `gorm:"primaryKey" json:"actorId"`
	CharacterID uint     `gorm:"primaryKey" json:"characterId"`
	AnimeID     uint     `gorm:"primaryKey" json:"animeId"`
	Role        CastRole `gorm:"type:cast_role;default:'supporting'" json:"role"`
}
