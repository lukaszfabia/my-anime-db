package models

// forms used in frontend

type LoginParams struct {
	Username string
	Password string
}

type Signup struct {
	Username string
	Email    string
	Password string
}

type UpdateAccount struct {
	Username string
	Email    string
	Password string
	Bio      string
	Website  string
	PicUrl   string
}
