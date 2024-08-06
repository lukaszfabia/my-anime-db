package models

type Response struct {
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

type TokenResponse struct {
	Response
	Token *string `json:"token,omitempty"`
}
