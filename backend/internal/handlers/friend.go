package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	friendcontroller "api/internal/controller/friend_controller"
	"api/internal/models"
	"api/pkg/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

var friendController controller.FriendController = &friendcontroller.FriendControllerImpl{}

func AddFriend(c *gin.Context) {
	friendId := c.Param("id") // id , user from ctx
	r := app.Gin{Ctx: c}
	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	if err := friendController.SendRequest(friendId, user); err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func DeleteFriend(c *gin.Context) {
	friendId := c.Param("id")
	r := app.Gin{Ctx: c}

	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	if err := friendController.DeleteRelation(user, friendId); err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func RespondToFriendRequest(c *gin.Context) {
	requestId := c.Param("id")
	r := app.Gin{Ctx: c}
	var action string = c.Query("status")

	status := tools.Match(models.AllFriendRequestStatus, action, models.Pending)

	if err := friendController.Respond(requestId, status); err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func GetInvitations(c *gin.Context) {
	user, err := controller.GetUserFromCtx(c)
	r := app.Gin{Ctx: c}
	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	pending, sent, err := friendController.GetRequests(user.ID)

	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, gin.H{
		"pending": pending,
		"sent":    sent,
	})
}

func GetFriendState(c *gin.Context) {
	senderId := c.Query("sender")
	receiverId := c.Query("receiver")
	r := app.Gin{Ctx: c}

	if request, err := friendController.GetState(senderId, receiverId); err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusOK, app.Ok, request)
	}

}
