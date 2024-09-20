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
	Website        string          `gorm:"default:''" json:"website,omitempty"`
	Posts          []Post          `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Friends        []*User         `gorm:"many2many:users_friends;" json:"friends,omitempty"`
	UserStats      *UserStat       `json:"stats,omitempty"`
	Reviews        []*Review       `gorm:"foreignKey:UserID" json:"reviews,omitempty"`
	FriendRequests []FriendRequest `gorm:"foreignKey:ReceiverID" json:"friendRequests,omitempty"`
}

type Post struct {
	gorm.Model
	Title    string  `gorm:"not null" json:"title"`
	Content  string  `gorm:"not null" json:"content"`
	Image    *string `json:"image,omitempty"`
	IsPublic bool    `json:"isPublic"`
	UserID   uint    `json:"userId"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

type FriendRequest struct {
	gorm.Model
	SenderID   uint                `gorm:"not null" json:"senderId"`
	ReceiverID uint                `gorm:"not null;column=receiver_id" json:"receiverId"`
	Status     FriendRequestStatus `gorm:"not null" json:"status"`

	Sender   User `gorm:"foreignKey:SenderID" json:"sender"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"receiver"`
}

// direct adding
type Studio struct {
	gorm.Model
	Name            string     `gorm:"not null;unique" json:"name" form:"name" binding:"required"`
	Website         string     `gorm:"default" json:"website,omitempty" form:"website"`
	EstablishedDate *time.Time `gorm:"type=date;" json:"establishedDate" form:"establishedDate" binding:"required"`
	LogoUrl         *string    `json:"logoUrl,omitempty" form:"pic"`

	Anime []*Anime `gorm:"foreignKey:StudioID" json:"animes"`
}

// direct adding
type Genre struct {
	gorm.Model
	Name string `gorm:"not null;unique;type=text" json:"name" binding:"required"`
}

type Anime struct {
	gorm.Model
	Title             string         `gorm:"not null;unique" json:"title" form:"title" binding:"required"`
	AlternativeTitles []*OtherTitles `gorm:"foreignKey:AnimeID" json:"alternativeTitles"`
	Type              AnimeType      `gorm:"type:text;default:'tv'" json:"animeType" form:"animeType"`
	Episodes          int            `gorm:"default:0" json:"episodes" form:"episodes"`
	Description       string         `gorm:"not null" json:"description" form:"description" binding:"required"`
	EpisodeLength     int            `gorm:"default:24" json:"episodeLength" form:"episodeLength"`
	StartDate         *time.Time     `gorm:"type:date,default" json:"startDate" form:"startDate"`
	FinishDate        *time.Time     `gorm:"type:date,default" json:"finishDate" form:"finishDate"`
	Pegi              Pegi           `gorm:"type:text;default:'PG-13'" json:"pegi" form:"pegi"`
	Status            StatusAnime    `gorm:"type:text;default:'unknown'" json:"status" form:"status"`
	PicUrl            *string        `json:"picUrl,omitempty" form:"pic"`
	Genres            []*Genre       `gorm:"many2many:anime_genres;" json:"genres" form:"genres"`
	StudioID          uint           `gorm:"not null;autoIncrement:false" form:"studio" json:"studioId" binding:"required"`
	PrequelID         *uint          `gorm:"column:prequel_id;autoIncrement:false" form:"prequelId" json:"prequelId,omitempty"`
	SequelID          *uint          `gorm:"column:sequel_id;autoIncrement:false" form:"sequelId" json:"sequelId,omitempty"`

	Studio    *Studio    `gorm:"foreignKey:StudioID" json:"studio"`
	Prequel   *Anime     `gorm:"foreignKey:PrequelID" json:"prequel,omitempty"`
	Sequel    *Anime     `gorm:"foreignKey:SequelID" json:"sequel,omitempty"`
	Roles     []*Role    `gorm:"foreignKey:AnimeID" json:"roles"`
	AnimeStat *AnimeStat `json:"stats"`
	Reviews   []*Review  `gorm:"foreignKey:AnimeID" json:"reviews,omitempty"`
}

// indirect adding through Anime
type OtherTitles struct {
	gorm.Model
	AnimeID          uint   `gorm:"not null;autoIncrement:false" json:"animeId"`
	AlternativeTitle string `gorm:"not null;unique" json:"title"`
}

// direct adding
type VoiceActor struct {
	gorm.Model
	Name      string  `gorm:"not null" json:"name" form:"name" binding:"required"`
	LastName  string  `gorm:"not null" json:"lastname" form:"lastname" binding:"required"`
	PicUrl    *string `json:"picUrl,omitempty" form:"picUrl"`
	Birthdate string  `gorm:"type:date" json:"birthdate" form:"birthdate"`
	Roles     []Role  `gorm:"foreignKey:ActorID" json:"roles,omitempty"`
}

// direct adding
type Character struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name" form:"name" binding:"required"`
	LastName    string  `gorm:"not null" json:"lastname" form:"lastname" binding:"required"`
	PicUrl      *string `json:"picUrl,omitempty" form:"picUrl"`
	Information string  `json:"information" form:"information" binding:"required"`
	Roles       []Role  `gorm:"foreignKey:CharacterID" json:"roles,omitempty"`
}

type UserAnime struct {
	gorm.Model
	UserID  uint        `gorm:"primaryKey;autoIncrement:false" json:"userId"`
	AnimeID uint        `gorm:"primaryKey;autoIncrement:false" json:"animeId"`
	Score   Score       `gorm:"type:text;default:'good'" json:"score"`
	Status  WatchStatus `gorm:"type:text;default:'plan-to-watch'" json:"watchStatus"`
	IsFav   bool        `gorm:"default:false" json:"isFav"`

	Anime Anime `gorm:"foreignKey:AnimeID" json:"anime"`
	User  User  `gorm:"foreignKey:UserID" json:"user"`
}

type Review struct {
	gorm.Model
	UserID  uint   `gorm:"not null" json:"userId"`
	AnimeID uint   `gorm:"not null" json:"animeId"`
	Content string `gorm:"not null" json:"content" form:"content" binding:"required"`

	UserAnime *UserAnime `gorm:"foreignKey:UserID,AnimeID;references:UserID,AnimeID" json:"userAnime"`
}

type Role struct {
	ActorID     uint     `gorm:"primaryKey;autoIncrement:false" json:"actorId"`
	CharacterID uint     `gorm:"primaryKey;autoIncrement:false" json:"characterId"`
	AnimeID     uint     `gorm:"primaryKey;autoIncrement:false" json:"animeId"`
	Role        CastRole `gorm:"type:text;default:'supporting'" json:"role"`

	VoiceActor VoiceActor `gorm:"foreignKey:ActorID;references:ID" json:"voiceActor"`
	Character  Character  `gorm:"foreignKey:CharacterID;references:ID" json:"character"`
	Anime      Anime      `gorm:"foreignKey:AnimeID;references:ID" json:"anime"`
}
