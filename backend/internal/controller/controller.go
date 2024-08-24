package controller

import (
	"api/internal/models"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FriendController interface {
	SendRequest(id string, user models.User) error
	DeleteRelation(user models.User, friendId string) error
	GetRequests(userId uint) ([]models.FriendRequest, []models.FriendRequest, error)
	GetState(senderId, receiverId string) (models.FriendRequest, error)
	Respond(requestId string, status models.FriendRequestStatus) error
}

type Controller[T models.Controllable] interface {
	GetAll() ([]T, error)
	Get(id string) (T, error)
	Create(c *gin.Context) (*T, error)
	Update(c *gin.Context, id string) (T, error)
	Delete(id string) error
}

func GetOrDefault(s string, def interface{}) interface{} {
	if s == "" {
		return def
	}

	switch def.(type) {
	case int:
		if res, err := strconv.Atoi(s); err == nil {
			return res
		}
	case float64:
		if res, err := strconv.ParseFloat(s, 64); err == nil {
			return res
		}
	case string:
		if strings.TrimSpace(s) != "" {
			return s
		}
	}

	return def
}

func GetUserFromCtx(c *gin.Context) (models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user.(models.User), nil
}
