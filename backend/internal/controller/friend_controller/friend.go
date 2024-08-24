package friendcontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"errors"
	"log"
)

type FriendControllerImpl struct{}

func (fc *FriendControllerImpl) SendRequest(id string, user models.User) error {
	var friend models.User

	if err := db.DB.First(&friend, id).Error; err != nil {
		return errors.New("user does not exists")
	}

	if user.ID == friend.ID {
		return errors.New("you cannot add yourself")
	}

	if res := db.DB.Model(&user).Association("Friends").Find(&friend); res != nil {
		return errors.New("you are already friends")
	}

	var existingRequest models.FriendRequest

	if err := db.DB.Where("sender_id = ? AND receiver_id = ?", user.ID, friend.ID).First(&existingRequest).Error; err == nil {
		return errors.New("you have already sent invite")
	}

	var newRequest models.FriendRequest = models.FriendRequest{
		SenderID:   user.ID,
		ReceiverID: friend.ID,
		Status:     models.Pending,
	}

	if err := db.DB.Create(&newRequest).Error; err != nil {
		return errors.New("couldn't create invite")
	}
	return nil
}

func (fc *FriendControllerImpl) DeleteRelation(user models.User, friendId string) error {
	var friend models.User

	if db.DB.First(friend, friendId).Error != nil {
		return errors.New("user does not exists")
	}

	if user.ID == friend.ID {
		return errors.New("can't remove yourself")
	}

	// removing friends from both sides
	if err := addOrDelete(user, friend, true); err != nil {
		return errors.New("couldn't remove friend")
	}

	return nil
}

func (fc *FriendControllerImpl) GetRequests(userId uint) ([]models.FriendRequest, []models.FriendRequest, error) {
	var invitations []models.FriendRequest
	var pendingInvitations []models.FriendRequest

	msgErr := "error during getting invitations"

	if err := db.DB.Preload("Sender").Where("receiver_id = ? AND status = ?", userId, models.Pending).
		Order("created_at DESC").Find(&invitations).Error; err != nil {
		return []models.FriendRequest{}, []models.FriendRequest{}, errors.New(msgErr)
	}

	if err := db.DB.Preload("Receiver").Where("sender_id = ? AND status = ?", userId, models.Pending).
		Order("created_at DESC").Find(&pendingInvitations).Error; err != nil {
		return []models.FriendRequest{}, []models.FriendRequest{}, errors.New(msgErr)
	}

	return invitations, pendingInvitations, nil
}

func (fc *FriendControllerImpl) GetState(senderId, receiverId string) (models.FriendRequest, error) {
	var request models.FriendRequest

	if err := db.DB.
		Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).
		Or("sender_id = ? AND receiver_id = ?", receiverId, senderId).
		Last(&request).Error; err != nil {
		return models.FriendRequest{}, errors.New("no requests")
	}

	return request, nil
}

func (fc *FriendControllerImpl) Respond(requestId string, status models.FriendRequestStatus) error {

	var request models.FriendRequest
	if res := db.DB.Where("id = ?", requestId).First(&request); res.Error != nil {
		return errors.New("couldn't find request")
	}

	var sender, receiver models.User
	senderRes := db.DB.First(&sender, request.SenderID)
	receiverRes := db.DB.First(&receiver, request.ReceiverID)

	if senderRes.Error != nil || receiverRes.Error != nil {
		return errors.New("sender or receiver id has not been found")
	}

	if status == models.Cancel {
		if err := cancelFriendRequest(sender.ID, receiver.ID); err != nil {
			return errors.New("couldn't cancel request")
		}

		return nil
	}

	// change status
	request.Status = status

	if err := db.DB.Save(&request).Error; err != nil {
		return errors.New("can't save request")
	}

	if err := addOrDelete(sender, receiver); err != nil {
		return errors.New("can't add friend")
	}

	return nil
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
