package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingIn(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong body",
		})
	}
}

func Login(c *gin.Context) {

}
