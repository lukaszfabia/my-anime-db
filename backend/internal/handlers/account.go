package handlers

import (
	"api/internal/models"
	"api/internal/utils"
	"api/pkg/db"
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

	picUrl := utils.SaveImage(c, utils.SaverProps{
		Dir:         utils.Avatar,
		Placeholder: utils.DefaultImage,
		KeyToImg:    "picUrl",
		Filename:    body.Username,
	})

	log.Println(picUrl)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
		PicUrl:   picUrl,
	}

	res := db.DB.Create(&user)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created !!",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}

func Me(c *gin.Context) {

	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "user does not exists",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
