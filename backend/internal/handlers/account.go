package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"api/pkg/utils"
	"api/pkg/validators"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SingUp handles the signup functionality for the API.
// It expects a POST request with the following form data:
// - username (string): the username of the user
// - password (string): the password of the user
// - email (string): the email address of the user
//
// It validates the form data and checks if the password is secure, the email has the correct structure,
// then hashes password, save avatar and object to database.
func SingUp(c *gin.Context) {
	var body models.SignupForm
	if !validators.IsFormDataValid(c, &body) {
		c.JSON(http.StatusBadRequest, response.BadForm())
		return
	}

	body = models.SignupForm{
		BaseForm: models.BaseForm{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		},
		Email: c.PostForm("email"),
	}

	if !validators.IsSecurePassword(body.Password) {
		msgErr := "Password is not secure"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	if !validators.IsEmailValid(body.Email) {
		msgErr := "Email has wrong structure"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	if !validators.IsUsernameValid(body.Username) {
		msgErr := "Username contains invalid characters"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	hash, err := utils.HashPassword(body.Password)

	if err != nil {
		msgErr := "Failed to hash password"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	picUrl := utils.SaveImage(c, utils.SaverProps{
		Dir:         utils.Avatar,
		Placeholder: utils.DefaultImage,
		KeyToImg:    "pic",
		Filename:    body.Username,
	})

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
		PicUrl:   &picUrl,
	}

	res := db.DB.Create(&user)

	if res.Error != nil {
		msgErr := "Failed to create user"
		c.JSON(http.StatusInternalServerError, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		utils.RemoveImage(*user.PicUrl)

		return

	}

	msg := "User has been created!!"
	c.JSON(http.StatusCreated, models.Response{
		Message: &msg,
		Error:   nil,
	})
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
	var body models.LoginForm

	if !validators.IsFormDataValid(c, &body) {
		c.JSON(http.StatusBadRequest, response.BadForm())
		return
	}

	body = models.LoginForm{
		BaseForm: models.BaseForm{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		},
	}

	var user models.User

	db.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		msgErr := "User not found"
		c.JSON(http.StatusNotFound, response.GenerateTokenResponse(nil, &msgErr, nil))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		msgErr := "Wrong password"
		c.JSON(http.StatusNotFound, response.GenerateTokenResponse(nil, &msgErr, nil))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		msgErr := "Failed to generate token"
		c.JSON(http.StatusBadRequest, response.GenerateTokenResponse(nil, &msgErr, nil))

		return
	}

	msg := "Token generated"
	c.JSON(http.StatusOK, response.GenerateTokenResponse(&msg, nil, &tokenStr))
}

// Me handles getting basic informations from account for the API
// It expects a GET request
func Me(c *gin.Context) {

	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)
	var userDetails models.User
	if err := db.DB.Model(&models.User{}).
		Preload("Friends", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "pic_url", "is_verified", "created_at", "bio", "website")
		}).
		Preload("Posts").
		Preload("UserAnimes").
		First(&userDetails, user.ID).Error; err != nil {
		// msgErr := "User not found"
		msgErr := err.Error()
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, userDetails)
}

// DeleteMe handles remove account functionality, bases on abstract function to deleting
// It expects a DELETE request
func DeleteMe(c *gin.Context) {

	user, _ := c.Get("user")

	castedUser, ok := user.(models.User)

	if !ok {
		msgErr := "Invalid user type"

		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
		return
	}

	res := db.Delete(models.User{}, strconv.Itoa(int(castedUser.ID)), "UserAnimes", "Posts", "Friends")
	if res != nil {
		msgErr := strings.ToTitle(res.Error())
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Account has been deleted"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))

}

// EditMe is a handler function that handles the editing of user account information.
// It expects a POST request with form data containing the updated account information.
// Validates form, gets user from context to get ID and retrieves user from database, then
// creates temporary object to store data from form a saves data.
func EditMe(c *gin.Context) {

	var body models.UpdateAccountForm

	if !validators.IsFormDataValid(c, &body) {
		c.JSON(http.StatusBadRequest, response.BadForm())
	}

	userCtx, _ := c.Get("user")
	user, ok := userCtx.(models.User)

	if !ok {
		msgErr := "Can't cast current user to model"
		c.JSON(http.StatusInternalServerError, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	pic, _ := c.FormFile("pic")

	body = models.UpdateAccountForm{
		Username: c.DefaultPostForm("username", user.Username),
		Password: c.PostForm("password"),
		Email:    c.DefaultPostForm("email", user.Email),
		Bio:      c.DefaultPostForm("bio", user.Bio),
		Website:  c.DefaultPostForm("webstie", *user.Website),
		PicFile:  pic,
	}

	db.DB.First(&user, user.ID)

	user.Username = body.Username
	user.Email = body.Email
	user.Bio = body.Bio
	*user.Website = body.Website

	// if there was a file
	if pic != nil {
		utils.RemoveImage(*user.PicUrl)
		*user.PicUrl = utils.SaveImage(c, utils.SaverProps{
			Filename: user.Username,
			KeyToImg: "picUrl",
		})
	}

	if newPassword := body.Password; validators.IsSecurePassword(newPassword) {
		hashedNewPass, err := utils.HashPassword(newPassword)
		if err == nil {
			user.Password = string(hashedNewPass)
		}
	}

	db.DB.Save(&user)

	msg := "User updated successfully!"

	c.JSON(http.StatusOK, models.Response{
		Message: &msg,
		Error:   nil,
	})

}
