package routes

import (
	"api/internal/handlers"
	"api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.Engine) {

	router.Static("/styles", "./styles")

	router.GET("/", handlers.Home)

	api := router.Group("/api")
	{
		api.POST("/sign-in", handlers.SingIn)
		api.POST("/login", handlers.Login)
		api.POST("/logout", handlers.Logout)

		auth := api.Group("/auth")
		{
			account := auth.Group("/account")
			{
				account.GET("/me", middleware.RequireAuth, handlers.Me)
			}

			// anime := auth.Group("/anime")
			// {

			// }
		}
	}

}
