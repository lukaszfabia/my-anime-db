package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	accountcontroller "api/internal/controller/account_controller"
	"api/internal/models"
	"api/pkg/validators"
	accountvalidator "api/pkg/validators/account_validator"
	loginvalidator "api/pkg/validators/login_validator" // Add this line

	"net/http"

	"github.com/gin-gonic/gin"
)

var accountController accountcontroller.AccountController = &accountcontroller.AccountControllerImpl{}
var accv validators.Validator = &accountvalidator.AccountValidator{}

// SingUp handles the signup functionality for the API.
// It expects a POST request with the following form data:
// - username (string): the username of the user
// - password (string): the password of the user
// - email (string): the email address of the user
//
// It validates the form data and checks if the password is secure, the email has the correct structure,
// then hashes password, save avatar and object to database.
func SingUp(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if !accv.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if err := accountController.CreateAccount(c); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusCreated, app.Ok, nil)
}

// Login handles the login functionality for the API.
//
// It expects a POST request with the following form data:
// - username: The username of the user.
// - password: The password of the user.
//
// Validates form, grabs user from context, compares with password in database,
// generates token if password is correct.
func Login(c *gin.Context) {
	var lgValid validators.Validator = &loginvalidator.LoginValidator{}
	var body models.LoginForm
	r := app.Gin{Ctx: c}

	if !lgValid.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	body = models.LoginForm{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	token, err := accountController.GenerateToken(body)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, token)
}

// Me handles getting basic informations from account for the API
// It expects a GET request
func Me(c *gin.Context) {
	r := app.Gin{Ctx: c}
	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	result, err := accountController.Get(&user)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, result)
}

// DeleteMe handles remove account functionality, bases on abstract function to deleting
// It expects a DELETE request
func DeleteMe(c *gin.Context) {
	r := app.Gin{Ctx: c}
	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	if err := accountController.DeleteAccount(&user); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

// EditMe is a handler function that handles the editing of user account information.
// It expects a POST request with form data containing the updated account information.
// Validates form, gets user from context to get ID and retrieves user from database, then
// creates temporary object to store data from form a saves data.
func EditMe(c *gin.Context) {
	r := app.Gin{Ctx: c}
	user, err := controller.GetUserFromCtx(c)

	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	if !accv.Validate(c) {
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if err := accountController.EditAccount(c, &user); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func SendCode(c *gin.Context) {
	user, err := controller.GetUserFromCtx(c)
	r := app.Gin{Ctx: c}
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	if err := accountController.SendCode(&user); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func Verify(c *gin.Context) {
	user, err := controller.GetUserFromCtx(c)
	r := app.Gin{Ctx: c}
	if err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	if err := accountController.VerifyAccount(c, &user); err != nil {
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}
