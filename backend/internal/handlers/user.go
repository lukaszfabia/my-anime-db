package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	usercontroller "api/internal/controller/user_controller"
	"api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userController controller.Controller[models.User] = &usercontroller.UserController{}

// GetAllUsers retrieves all users from the database.
//
// This function queries the database to fetch all users and returns them as a JSON response.
// If there is an error during the database query, it returns a JSON response with an error message.
func GetAllUsers(c *gin.Context) {
	r := app.Gin{Ctx: c}

	users, err := userController.GetAll()
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, users)
}

// RetrieveUser retrieves a user from the database based on the provided ID.
//
// It returns the user information in JSON format if found, otherwise it returns an error message.
// The user's friends, posts, and user animes are also preloaded.
func GetUser(c *gin.Context) {
	id := c.Param("id")

	r := app.Gin{Ctx: c}

	user, err := userController.Get(id)
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, user)
}
