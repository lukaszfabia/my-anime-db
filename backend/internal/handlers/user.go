package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiUser struct {
	ShortApiUser
	CreatedAt  time.Time
	Bio        string              `json:"bio"`
	Website    string              `json:"website"`
	Friends    []*models.User      `gorm:"many2many:users_friends;" json:"friends"`
	Posts      []*models.Post      `gorm:"many2many:users_posts;" json:"posts"`
	UserAnimes []*models.UserAnime `gorm:"many2many:users_anime" json:"userAnimes"`
}

type ShortApiUser struct {
	Username   string `json:"username"`
	PicUrl     string `json:"picUrl"`
	IsVerified bool   `json:"isVerified"`
	IsMod      bool   `json:"isMod"`
}

func GetAllUsers(c *gin.Context) {
	var apiUsers []ShortApiUser

	db.DB.Model(&models.User{}).
		Select(&ShortApiUser{}).
		Where("username <> ?", "").
		Take(&apiUsers)

	c.JSON(http.StatusOK, apiUsers)
}

func RetriveUser(c *gin.Context) {
	id := c.Param("id")

	var apiUser ApiUser

	res := db.DB.
		Preload("Friends").
		Preload("Posts").
		Preload("UserAnimes").
		Model(&models.User{}).
		Where("id = ? AND username <> ''", id).
		First(&apiUser)

	if res.Error != nil {
		msgErr := "User not found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": apiUser,
	})

}
