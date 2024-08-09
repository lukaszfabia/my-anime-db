package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ShortReadOnlyUser struct {
	ID         uint    `json:"id"`
	Username   string  `json:"username"`
	PicUrl     *string `json:"picUrl"`
	IsVerified bool    `json:"isVerified"`
}

type ReadOnlyUser struct {
	ShortReadOnlyUser
	CreatedAt  time.Time
	Bio        string              `json:"bio,omitempty"`
	Website    string              `json:"website,omitempty"`
	Friends    []*models.User      `gorm:"many2many:users_friends;" json:"friends,omitempty"`
	Posts      []*models.Post      `gorm:"many2many:users_posts;" json:"posts,omitempty"`
	UserAnimes []*models.UserAnime `gorm:"many2many:users_anime" json:"userAnimes,omitempty"`
}

// GetAllUsers retrieves all users from the database.
//
// This function queries the database to fetch all users and returns them as a JSON response.
// If there is an error during the database query, it returns a JSON response with an error message.
func GetAllUsers(c *gin.Context) {
	var apiUsers []ShortReadOnlyUser

	res := db.DB.Model(&models.User{}).Find(&apiUsers)

	if res.Error != nil {
		msgErr := res.Error.Error()
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, apiUsers)
}

// RetrieveUser retrieves a user from the database based on the provided ID.
//
// It returns the user information in JSON format if found, otherwise it returns an error message.
// The user's friends, posts, and user animes are also preloaded.
func RetrieveUser(c *gin.Context) {
	id := c.Param("id")

	var apiUser ReadOnlyUser

	err := db.Retrieve(&models.User{}, &apiUser, id, "Posts", "Friends", "UserAnimes")

	if err != nil {
		msgErr := "User not found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": apiUser,
	})

}
