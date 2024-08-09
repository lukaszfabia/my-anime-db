package models

import "mime/multipart"

// forms used in frontend

type BaseForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginForm struct {
	BaseForm
}

type SignupForm struct {
	BaseForm
	Email string `form:"email" binding:"required"`
}

type UpdateAccountForm struct {
	Username string                `form:"username" json:"username,omitempty"`
	Email    string                `form:"email" json:"email,omitempty" `
	Password string                `form:"password" json:"password,omitempty"`
	Bio      string                `form:"bio" json:"bio,omitempty"`
	Website  string                `form:"website" json:"website,omitempty"`
	PicFile  *multipart.FileHeader `form:"pic" json:"picUrl,omitempty"`
}
