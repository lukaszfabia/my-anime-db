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

func (va *VoiceActor) GetPicUrl() *string {
	return va.PicUrl
}

func (c *Character) GetPicUrl() *string {
	return c.PicUrl
}
