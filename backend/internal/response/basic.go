package response

import "api/internal/models"

func NewResponse(msg *string, msgErr *string) models.Response {
	return models.Response{
		Message: msg,
		Error:   msgErr,
	}
}

func BadForm() models.Response {
	msgErr := "Invalid form"
	return models.Response{
		Message: nil,
		Error:   &msgErr,
	}
}
