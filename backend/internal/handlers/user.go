package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShortReadOnlyUser struct {
	ID         uint    `json:"id"`
	Username   string  `json:"username"`
	PicUrl     *string `json:"picUrl"`
	IsVerified bool    `json:"isVerified"`
}

// GetAllUsers retrieves all users from the database.
//
// This function queries the database to fetch all users and returns them as a JSON response.
// If there is an error during the database query, it returns a JSON response with an error message.
func GetAllUsers(c *gin.Context) {
	var apiUsers []ShortReadOnlyUser

	if err := db.RetrieveAll(&models.User{}, &apiUsers, db.ToOrder(db.DB, "username DESC")); err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
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

	var apiUser models.User

	err := db.DB.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Preload("Friends", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "pic_url", "is_verified", "created_at", "bio", "website")
	}).
		Preload("UserAnimes"). // todo add another select
		First(&apiUser, id).Error

	if err != nil {
		msgErr := "user not found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}
	c.JSON(http.StatusOK, apiUser)
}
