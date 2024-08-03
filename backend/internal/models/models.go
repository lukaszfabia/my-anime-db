package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string      `gorm:"unique;not null" json:"username"`
	Email      string      `gorm:"unique;not null" json:"email"`
	Password   string      `gorm:"not null" json:"password"`
	PicUrl     string      `gorm:"default:'https://placehold.co/400'" json:"picUrl"`
	IsVerified bool        `gorm:"default:false" json:"isVerified"`
	IsMod      bool        `gorm:"default:false" json:"isMod"`
	Bio        string      `gorm:"default:'Edit bio'" json:"bio"`
	Website    string      `json:"website"`
	Posts      []Post      `gorm:"many2many:users_posts;" json:"posts"`
	Friends    []User      `gorm:"many2many:users_friends;" json:"friends"`
	UserAnimes []UserAnime `gorm:"foreignKey:UserID" json:"userAnimes"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	Image    string
	IsPublic bool `gorm:"default:true"`
}

type Studio struct {
	gorm.Model
	Name    string `gorm:"not null;unique"`
	Website string
}

type Genre struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

type Anime struct {
	gorm.Model
	Title             string `gorm:"not null;unique"`
	AlternativeTitles string
	Type              AnimeType  `gorm:"type:anime_type;default:'tv'"`
	Episodes          int        `gorm:"default:0"`
	Descripting       string     `gorm:"not null"`
	EpisodeLength     int        `gorm:"default:24"`
	StartDate         *time.Time `json:"startDate"`
	FinishDate        *time.Time `json:"finishDate"`
	Pegi              Pegi       `gorm:"type:pegi;default:'PG-13'"`
	PicUrl            string     `gorm:"default:'https://placehold.co/400'" json:"picUrl"`
	Genres            []Genre    `gorm:"many2many:anime_gernes;"`
	Studios           []Studio   `gorm:"many2many:anime_studios;"`
}

type VoiceActor struct {
	gorm.Model
	Name      string `gorm:"not-null"`
	LastName  string `gorm:"not-null"`
	PicUrl    string `gorm:"default:'https://placehold.co/400'" json:"picUrl"`
	BirthDate *time.Time
}

type Character struct {
	gorm.Model
	Name        string `gorm:"not-null"`
	Information string `gorm:"default:'update info'"`
	PicUrl      string `gorm:"default:'https://placehold.co/400'" json:"picUrl"`
}

type UserAnime struct {
	UserID  uint   `gorm:"primaryKey"`
	AnimeID uint   `gorm:"primaryKey"`
	Score   Score  `gorm:"type:score;default:'good'"`
	Status  Status `gorm:"type:status;default:'plan-to-watch'"`
	Review  string
}

type Roles struct {
	ActorID     uint `gorm:"primaryKey"`
	CharacterID uint `gorm:"primaryKey"`
	AnimeID     uint `gorm:"primaryKey"`
	Role        Role `gorm:"type:role;default:'supporting'"`
}
