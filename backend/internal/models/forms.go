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
