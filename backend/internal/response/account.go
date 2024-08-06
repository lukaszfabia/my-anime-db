package response

import "api/internal/models"

func GenerateTokenResponse(msg *string, msgErr *string, token *string) models.TokenResponse {
	return models.TokenResponse{
		Response: NewResponse(msg, msgErr),
		Token:    token,
	}
}
