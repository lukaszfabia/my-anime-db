package models

import (
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
	UserAnimes     []*UserAnime    `gorm:"foreignKey:UserID" json:"userAnimes,omitempty"`
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
	Name    string `gorm:"not null;unique" json:"name" binding:"required"`
	Website string `json:"website"`
}

// direct adding
type Genre struct {
	gorm.Model
	Name GenreOption `gorm:"not null;unique;type=text" json:"name" binding:"required"`
}

type Anime struct {
	gorm.Model
	Title             string         `gorm:"not null;unique" json:"title" binding:"required"`
	AlternativeTitles []*OtherTitles `gorm:"many2many:other_titles;" json:"alternativeTitles"`
	Type              AnimeType      `gorm:"type:text;default:'tv'" json:"type"`
	Episodes          int            `gorm:"default:0" json:"episodes"`
	Description       string         `gorm:"not null" json:"description" binding:"required"`
	EpisodeLength     int            `gorm:"default:24" json:"episodeLength"`
	StartDate         string         `gorm:"type:date" json:"startDate"`
	FinishDate        string         `gorm:"type:date" json:"finishDate"`
	Pegi              Pegi           `gorm:"type:text;default:'PG-13'" json:"pegi"`
	Status            StatusAnime    `gorm:"type:text;default:'unknown'" json:"status"`
	PicUrl            *string        `json:"picUrl"`
	AnimeStat         *AnimeStat     `json:"stats"`
	Genres            []*Genre       `gorm:"many2many:anime_genres;" json:"genres"`
	Studio            *Studio        `gorm:"foreignKey:ID;" json:"studio"`
	Roles             []*Role        `gorm:"foreignKey:AnimeID" json:"roles"`
	Prequel           *Anime         `gorm:"foreignKey:ID" json:"prequel,omitempty"`
	Sequel            *Anime         `gorm:"foreignKey:ID" json:"sequel,omitempty"`
}

type AnimeStat struct {
	gorm.Model
	AnimeID     uint     `gorm:"primaryKey" json:"animeId"`
	GlobalScore *float64 `gorm:"default:0.0" json:"score"`
	Popularity  *uint    `gorm:"default:0" json:"popularity"`
}

// indirect adding through Anime
type OtherTitles struct {
	gorm.Model
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
	Information string  `json:"information" form:"information"`
	Roles       []Role  `gorm:"foreignKey:ActorID" json:"roles,omitempty"`
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

	// preoading depends on the used models e.g voice actor doesnt need to preload voice actor
	VoiceActor VoiceActor `gorm:"foreignKey:ActorID" json:"voiceActor"`
	Character  Character  `gorm:"foreignKey:CharacterID" json:"character"`
	Anime      Anime      `gorm:"foreignKey:AnimeID" json:"anime"`
}
