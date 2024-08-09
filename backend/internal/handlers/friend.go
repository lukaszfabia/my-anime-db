package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"api/pkg/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFriend(c *gin.Context) {
	friendId := c.Param("id")

	var friend models.User

	res := db.DB.Model(models.User{}).First(&friend, friendId)

	if res.Error != nil {
		msgErr := "User does not exists"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))

		return
	}

	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)

	var existingRequest models.FriendRequest

	if res := db.DB.Where("sender_id = ? AND reciever_id = ?", user.ID, friend.ID).First(&existingRequest); res.Error == nil {
		msgErr := "You have already sent invite"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	var newRequest models.FriendRequest = models.FriendRequest{
		SenderID:   user.ID,
		RecieverID: friend.ID,
		Status:     models.Peding,
	}

	if res := db.DB.Create(&newRequest); res.Error != nil {
		msgErr := "Couldn't create invite"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Request has been sent"
	c.JSON(http.StatusCreated, response.NewResponse(&msg, nil))
}

func RemoveFriend(c *gin.Context) {
	id := c.Param("id")
	var friend models.User

	db.DB.First(&friend, id)

	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)

	if err := db.DB.Model(&user).Association("Friends").Delete(&friend); err != nil {
		msgErr := "Could't remove from friends"
		c.JSON(http.StatusBadGateway, response.NewResponse(nil, &msgErr))
		return
	}
}

func RespondToFriendRequest(c *gin.Context) {
	requestId := c.Param("id")
	var action string = c.Query("status")

	status := tools.Match(models.AllFriendRequestStatus, action, models.Peding)

	var request models.FriendRequest
	if res := db.DB.Where("id = ?", requestId).First(&request); res.Error != nil {
		msgErr := "Couldn't find request"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	var sender, reciever models.User

	senderRes := db.DB.First(&sender, request.SenderID)
	recieverRes := db.DB.First(&reciever, request.RecieverID)

	if senderRes.Error != nil && recieverRes.Error != nil {
		msgErr := "Sender or reciever id has not been found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	request.Status = status

	if err := db.DB.Save(&request).Error; err != nil {
		msgErr := "Can't save request"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	if err := db.DB.Model(&sender).Association("Friends").Append(&reciever); err != nil {
		msgErr := "Couldn't save relation"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Successfully added to friends"
	c.JSON(http.StatusCreated, response.NewResponse(&msg, nil))
}

func GetInvitations(c *gin.Context) {
	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)

	var invitations []models.FriendRequest

	if err := db.DB.Where("reciever_id = ? AND status = ?", user.ID, models.Peding).Find(&invitations).Error; err != nil {
		msgErr := "Couldn't find any invitations"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, invitations)
}
