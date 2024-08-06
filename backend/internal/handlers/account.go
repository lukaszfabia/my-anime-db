package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/internal/utils"
	"api/pkg/db"
	"api/pkg/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SingUp(c *gin.Context) {

	var body models.Signup = models.Signup{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if !middleware.IsSecurePassword(body.Password) {
		msgErr := "Password is not secure"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	if !middleware.IsEmailValid(body.Email) {
		msgErr := "Email has wrong structure"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		return
	}

	if !middleware.IsUsernameValid(body.Username) {
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

	log.Println(picUrl)

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
		PicUrl:   picUrl,
	}

	res := db.DB.Create(&user)

	if res.Error != nil {
		msgErr := "Failed to create user"
		c.JSON(http.StatusInternalServerError, models.Response{
			Error:   &msgErr,
			Message: nil,
		})

		os.Remove(picUrl)

		return

	}

	msg := "User has been created!!"
	c.JSON(http.StatusCreated, models.Response{
		Message: &msg,
		Error:   nil,
	})
}

func Login(c *gin.Context) {

	var body models.LoginParams = models.LoginParams{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
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

func Me(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, user)
}

func DeleteMe(c *gin.Context) {

	user, _ := c.Get("user")

	castedUser, ok := user.(models.User)

	if !ok {
		msgErr := "Invalid user type"
		c.JSON(http.StatusInternalServerError, models.Response{Error: &msgErr, Message: nil})
		return
	}

	avatarPath := castedUser.PicUrl

	log.Println(castedUser.ID)

	res := db.DB.Unscoped().Delete(&user)

	if res.Error != nil {
		log.Printf("Error deleting user: %v", res.Error)
		msgErr := "Can't remove account"
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   &msgErr,
			Message: nil,
		})
		return
	}

	utils.RemoveImage(avatarPath)

	msg := "Account has been deleted"
	c.JSON(http.StatusOK, models.Response{
		Message: &msg,
		Error:   nil,
	})

}

func EditMe(c *gin.Context) {

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

	pic, _ := c.FormFile("picUrl")

	db.DB.Where("id = ?", user.ID).First(&user)

	user.Username = c.DefaultPostForm("username", user.Username)

	if pic != nil {
		utils.RemoveImage(user.PicUrl)
		user.PicUrl = utils.SaveImage(c, utils.SaverProps{
			Filename: user.Username,
			KeyToImg: "picUrl",
		})
	}

	if newPassword := c.PostForm("password"); newPassword != "" && middleware.IsSecurePassword(newPassword) {
		hashedNewPass, err := utils.HashPassword(newPassword)
		if err == nil {
			user.Password = string(hashedNewPass)
		}
	}

	user.Email = c.DefaultPostForm("email", user.Email)
	user.Bio = c.DefaultPostForm("bio", user.Bio)
	user.Website = c.DefaultPostForm("website", user.Website)

	db.DB.Save(&user)

	msg := "User updated successfully!"

	c.JSON(http.StatusOK, models.Response{
		Message: &msg,
		Error:   nil,
	})

}
