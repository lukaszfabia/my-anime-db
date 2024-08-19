package models

type PicUrlGetter interface {
	GetPicUrl() *string
}

func (a *Anime) GetPicUrl() *string {
	return a.PicUrl
}

func (u *User) GetPicUrl() *string {
	return u.PicUrl
}

func (p *Post) GetPicUrl() *string {
	return p.Image
}
