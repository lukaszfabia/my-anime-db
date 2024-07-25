package routes

import (
	"api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/sign-in", handlers.SingIn)
	}
}
