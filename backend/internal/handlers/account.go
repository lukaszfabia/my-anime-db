package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/internal/store"
	"api/pkg/db"
	"api/pkg/tools"
	"api/pkg/utils"
	"api/pkg/validators"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var storage = store.NewVerificationStore()

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
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
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

	userObj, _ := c.Get("user")

	user, _ := userObj.(models.User)

	res := db.Delete(models.User{}, strconv.Itoa(int(user.ID)),
		db.Association{Model: "UserAnimes"},
		db.Association{Model: "Posts"},
		db.Association{Model: "Friends"},
	)

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

	user, _ := userCtx.(models.User)

	pic, _ := c.FormFile("pic")

	if err := db.DB.First(&user, user.ID).Error; err != nil {
		msgErr := "User not found"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	body = models.UpdateAccountForm{
		Username: tools.GetOrDefault(c.PostForm("username"), user.Username).(string),
		Password: c.PostForm("password"),
		Email:    tools.GetOrDefault(c.PostForm("email"), user.Email).(string),
		Bio:      tools.GetOrDefault(c.PostForm("bio"), user.Bio).(string),
		Website:  tools.GetOrDefault(c.PostForm("website"), user.Website).(string),
		PicFile:  pic,
	}

	user.Username = body.Username

	log.Println(body.Email, user.Email)

	if body.Email != user.Email {
		user.IsVerified = false
	}

	user.Email = body.Email

	user.Bio = body.Bio
	user.Website = body.Website

	// if there was a file
	if pic != nil {
		utils.RemoveImage(*user.PicUrl)
		*user.PicUrl = utils.SaveImage(c, utils.SaverProps{
			Filename: user.Username,
			KeyToImg: "pic",
		})
	}

	if newPassword := body.Password; validators.IsSecurePassword(newPassword) {
		hashedNewPass, err := utils.HashPassword(newPassword)
		if err == nil {
			user.Password = string(hashedNewPass)
		}
	}

	if err := db.DB.Save(&user).Error; err != nil {
		msgErr := "Cannot update user"
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "User updated successfully!"

	c.JSON(http.StatusOK, models.Response{
		Message: &msg,
		Error:   nil,
	})

}

func SendCode(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user, _ := userCtx.(models.User)

	senderMail := os.Getenv("GOOGLE_MAIL")
	senderPassword := os.Getenv("GOOGLE_PASSWORD")

	auth := smtp.PlainAuth("", senderMail, senderPassword, "smtp.gmail.com")

	to := []string{user.Email}

	code := utils.GenerateCode()

	storage.Set(user.Email, code)

	message := []byte(fmt.Sprintf(
		"To: %s\r\n"+

			"Subject: Account verification\r\n"+

			"\r\n"+

			"Hello %s.\r\n"+

			"Here is your pin: %s.\r\n"+

			"Note: Please enter it in 2 minutes. After that your code expires.", user.Email, user.Username, code))

	err := smtp.SendMail("smtp.gmail.com:587", auth, senderMail, to, message)

	if err != nil {
		msgErr := "Failed to send email"
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Email has been sent"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}

func Verify(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user, _ := userCtx.(models.User)

	code := c.PostForm("code")

	if err := storage.Compare(code, user.Email); err != nil {
		msgErr := "Wrong code"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	user.IsVerified = true

	if err := db.DB.Save(&user).Error; err != nil {
		msgErr := "Cannot update user"
		c.JSON(http.StatusInternalServerError, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "Account has been verifed successfully!"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}
