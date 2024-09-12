package models

import "mime/multipart"

// forms used in frontend

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupForm struct {
	LoginForm
	Email string `form:"email" json:"email" binding:"required"`
}

type UpdateAccountForm struct {
	Username string                `form:"username" json:"username,omitempty"`
	Email    string                `form:"email" json:"email,omitempty" `
	Password string                `form:"password" json:"password,omitempty"`
	Bio      string                `form:"bio" json:"bio,omitempty"`
	Website  string                `form:"website" json:"website,omitempty"`
	PicFile  *multipart.FileHeader `form:"pic" json:"picUrl,omitempty"`
}

type PostForm struct {
	Title    string `form:"title" json:"title" binding:"required"`
	IsPublic bool   `form:"isPublic" json:"isPublic"`
	Content  string `form:"content" json:"content" binding:"required"`
}

type AnimeForm struct {
	Title             string   `form:"title" json:"title" binding:"required"`
	AlternativeTitles []string `form:"altTitles" json:"altTitles"`
	Type              string   `form:"animeType" json:"type" binding:"required"`
	Episodes          string   `form:"episodes" json:"episodes"`
	Description       string   `form:"description" json:"description" binding:"required"`
	EpisodeLength     string   `form:"episodeLength" json:"episodeLength"`
	StartDate         string   `form:"startDate" json:"startDate"`
	FinishDate        string   `form:"finishDate" json:"finishDate"`
	Pegi              string   `form:"pegi" json:"pegi"`
	Status            string   `form:"status" json:"status"`
	Genres            []string `form:"genres" json:"genres"`
	StudioID          string   `form:"studio" json:"studioId" binding:"required"`
	PrequelID         string   `form:"prequelId" json:"prequelId,omitempty"`
	SequelID          string   `form:"sequelId" json:"sequelId,omitempty"`
}
