package models

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Signup struct {
	Username string `json:"username" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
