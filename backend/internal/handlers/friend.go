package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"api/pkg/tools"
	"errors"
	"log"
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

	if user.ID == friend.ID {
		msgErr := "You can't add yourself"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	if res := db.DB.Model(&user).Association("Friends").Find(&friend); res != nil {
		msgErr := "You are already friends"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	var existingRequest models.FriendRequest

	if err := db.DB.Where("sender_id = ? AND receiver_id = ?", user.ID, friend.ID).First(&existingRequest).Error; err == nil {
		msgErr := "You have already sent invite"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	var newRequest models.FriendRequest = models.FriendRequest{
		SenderID:   user.ID,
		ReceiverID: friend.ID,
		Status:     models.Pending,
	}

	if err := db.DB.Create(&newRequest).Error; err != nil {
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

	if user.ID == friend.ID {
		msgErr := "You can't remove yourself"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	// removing friends from both sides
	if err := addOrDelete(user, friend, true); err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Friend has been removed"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}

func RespondToFriendRequest(c *gin.Context) {
	requestId := c.Param("id")
	var action string = c.Query("status")

	status := tools.Match(models.AllFriendRequestStatus, action, models.Pending)

	var request models.FriendRequest
	if res := db.DB.Where("id = ?", requestId).First(&request); res.Error != nil {
		msgErr := "couldn't find request"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	var sender, receiver models.User
	senderRes := db.DB.First(&sender, request.SenderID)
	receiverRes := db.DB.First(&receiver, request.ReceiverID)

	if senderRes.Error != nil || receiverRes.Error != nil {
		msgErr := "sender or receiver id has not been found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	if status == models.Cancel {
		if err := cancelFriendRequest(sender.ID, receiver.ID); err != nil {
			msgErr := err.Error()
			c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
			return
		}

		msg := "request has been canceled"
		c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
		return
	}
	// create add method
	request.Status = status

	if err := db.DB.Save(&request).Error; err != nil {
		msgErr := "can't save request"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	if err := addOrDelete(sender, receiver); err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "successfully added to friends"
	c.JSON(http.StatusCreated, response.NewResponse(&msg, nil))
}

func GetInvitations(c *gin.Context) {
	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)

	var invitations []models.FriendRequest
	var pendingInvitations []models.FriendRequest

	msgErr := "error during getting invitations"

	if err := db.DB.Preload("Sender").Where("receiver_id = ? AND status = ?", user.ID, models.Pending).
		Order("created_at DESC").Find(&invitations).Error; err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	if err := db.DB.Preload("Receiver").Where("sender_id = ? AND status = ?", user.ID, models.Pending).
		Order("created_at DESC").Find(&pendingInvitations).Error; err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invitations":        invitations,
		"pendingInvitations": pendingInvitations,
	})
}

func GetFriendState(c *gin.Context) {
	senderId := c.Query("sender")
	receiverId := c.Query("receiver")

	var request models.FriendRequest

	if err := db.DB.
		Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).
		Or("sender_id = ? AND receiver_id = ?", receiverId, senderId).
		Last(&request).Error; err != nil {
		msgErr := "no requests"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, request)
}

func cancelFriendRequest(senderId, receiverId uint) error {

	var request models.FriendRequest

	if err := db.DB.Where("sender_id = ? AND receiver_id = ? AND status = ?", senderId, receiverId, models.Pending).First(&request).Error; err != nil {
		return errors.New("no requests")
	}

	if err := db.DB.Unscoped().Delete(&request).Error; err != nil {
		return errors.New("couldn't delete request")
	}

	return nil
}

func addOrDelete(lhs, rhs models.User, onDelete ...bool) error {
	if len(onDelete) > 0 {
		log.Println("deleting")
		if err := db.DB.Model(&lhs).Association("Friends").Delete(&rhs); err != nil {
			return errors.New("couldn't save relation")
		}

		if err := db.DB.Model(&rhs).Association("Friends").Delete(&lhs); err != nil {
			return errors.New("couldn't save relation")
		}

		if err := db.DB.Model(&models.FriendRequest{}).
			Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", lhs.ID, rhs.ID, rhs.ID, lhs.ID).
			Unscoped().Delete(models.FriendRequest{}).Error; err != nil {
			return errors.New("couldn't remove friend")
		}

		return nil
	}

	if err := db.DB.Model(&lhs).Association("Friends").Append(&rhs); err != nil {
		return errors.New("couldn't save relation")
	}

	if err := db.DB.Model(&rhs).Association("Friends").Append(&lhs); err != nil {
		return errors.New("couldn't save relation")
	}

	return nil
}
