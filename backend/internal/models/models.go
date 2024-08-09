package models

import (
	"time"

	"gorm.io/gorm"
)

// orm

type User struct {
	gorm.Model
	Username       string          `gorm:"unique;not null" json:"username"`
	Email          string          `gorm:"unique;not null" json:"email"`
	Password       string          `gorm:"not null" json:"password"`
	PicUrl         *string         `json:"picUrl,omitempty"`
	IsVerified     bool            `gorm:"default:false" json:"isVerified"`
	IsMod          bool            `gorm:"default:false" json:"isMod"`
	Bio            string          `gorm:"default:'Edit bio'" json:"bio,omitempty"`
	Website        *string         `json:"website,omitempty"`
	Posts          []*Post         `gorm:"many2many:users_posts;" json:"posts,omitempty"`
	Friends        []*User         `gorm:"many2many:users_friends;" json:"friends,omitempty"`
	UserAnimes     []*UserAnime    `gorm:"foreignKey:UserID" json:"userAnimes,omitempty"`
	FriendRequests []FriendRequest `gorm:"foreignKey:RecieverID" json:"friendRequests,omitempty"`
}

type FriendRequest struct {
	gorm.Model
	SenderID   uint                `gorm:"not null" json:"sender"`
	RecieverID uint                `gorm:"not null" json:"reciever"`
	Status     FriendRequestStatus `gorm:"not null" json:"status"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" binding:"required" form:"title"`
	Content  string `gorm:"not null" json:"content" binding:"required" form:"content"`
	Image    string `json:"image" form:"image"`
	IsPublic bool   `gorm:"default:true" json:"isPublic" form:"isPublic"`
}

type Studio struct {
	gorm.Model
	Name    string `gorm:"not null;unique" json:"name" binding:"required"`
	Website string `json:"website"`
}

type Genre struct {
	gorm.Model
	Name GenreOption `gorm:"not null;unique;type=text" json:"name" binding:"required"`
}

type Anime struct {
	gorm.Model
	Title             string    `gorm:"not null;unique" json:"title" binding:"required"`
	AlternativeTitles string    `json:"alternativeTitles"`
	Type              AnimeType `gorm:"type:text;default:'tv'" json:"type"`
	Episodes          int       `gorm:"default:0" json:"episodes"`
	Description       string    `gorm:"not null" json:"description" binding:"required"`
	EpisodeLength     int       `gorm:"default:24" json:"episodeLength"`
	StartDate         string    `gorm:"type:date" json:"startDate"`
	FinishDate        string    `gorm:"type:date" json:"finishDate"`
	Pegi              Pegi      `gorm:"type:text;default:'PG-13'" json:"pegi"`
	PicUrl            *string   `json:"picUrl"`
	GlobalScore       *float64  `gorm:"default:0.0" json:"score"`
	Popularity        *uint     `gorm:"default:0" json:"popularity"`
	Genres            []*Genre  `gorm:"many2many:anime_genres;" json:"genres"`
	Studios           []*Studio `gorm:"many2many:anime_studios;" json:"studios"`
	Roles             []*Role   `gorm:"foreignKey:AnimeID" json:"roles"`
}

type VoiceActor struct {
	gorm.Model
	Name      string     `gorm:"not null" json:"name"`
	LastName  string     `gorm:"not null" json:"lastName"`
	PicUrl    string     `json:"picUrl"`
	BirthDate *time.Time `json:"birthDate"`
	Roles     []*Role    `gorm:"foreignKey:ActorID" json:"roles"`
}

type Character struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Information string  `gorm:"default:'update info'" json:"information"`
	PicUrl      string  `json:"picUrl"`
	Roles       []*Role `gorm:"foreignKey:CharacterID" json:"roles"`
}

type UserAnime struct {
	gorm.Model
	UserID  uint   `gorm:"primaryKey" json:"userId"`
	AnimeID uint   `gorm:"primaryKey" json:"animeId"`
	Score   Score  `gorm:"type:text;default:'good'" json:"score"`
	Status  Status `gorm:"type:text;default:'plan-to-watch'" json:"status"`
	Review  string `json:"review"`
}

type Role struct {
	gorm.Model
	ActorID     uint     `gorm:"primaryKey" json:"actorId"`
	CharacterID uint     `gorm:"primaryKey" json:"characterId"`
	AnimeID     uint     `gorm:"primaryKey" json:"animeId"`
	Role        CastRole `gorm:"type:text;default:'supporting'" json:"role"`
}
