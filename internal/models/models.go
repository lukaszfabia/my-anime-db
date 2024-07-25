package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Pic        string `gorm:"default:'https://placehold.co/400'"`
	IsVerified bool   `gorm:"default:false"`
	IsMod      bool   `gorm:"default:false"`
	Bio        string `gorm:"default:'Edit bio'"`
	Website    string
	Posts      []Post      `gorm:"many2many:users_posts;"`
	Friends    []User      `gorm:"many2many:users_friends;"`
	UserAnimes []UserAnime `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	Image    string `gorm:"default:'https://placehold.co/400x600'"`
	IsPublic bool   `gorm:"default:true"`
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
	Type              AnimeType `gorm:"type:anime_type;default:'tv'"`
	Episodes          int       `gorm:"default:0"`
	Descripting       string    `gorm:"not null"`
	EpisodeLength     int       `gorm:"default:24"`
	DateStarted       *time.Time
	DateFinished      *time.Time
	Pegi              Pegi     `gorm:"type:pegi;default:'PG-13'"`
	Pic               string   `gorm:"default:'https://placehold.co/400x600'"`
	Genres            []Genre  `gorm:"many2many:anime_gernes;"`
	Studios           []Studio `gorm:"many2many:anime_studios;"`
}

type VoiceActor struct {
	gorm.Model
	Name      string `gorm:"not-null"`
	LastName  string `gorm:"not-null"`
	Pic       string `gorm:"default:'https://placehold.co/400x600'"`
	BirthDate *time.Time
}

type Character struct {
	gorm.Model
	Name        string `gorm:"not-null"`
	Information string `gorm:"default:'update info'"`
	Pic         string `gorm:"default:'https://placehold.co/400x600'"`
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
