package routes

import (
	"api/internal/handlers"
	"api/templates"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.Engine) {

	elem := templates.Base()

	router.Static("/static", "./static")
	router.Static("/styles", "./styles")

	router.GET("/", func(c *gin.Context) {
		templ.Handler(elem).ServeHTTP(c.Writer, c.Request)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/sign-in", handlers.SingIn)
	}
}
