package routes

import (
	"api/internal/handlers"
	"api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.Engine) {
	router.GET("/", handlers.Home)

	router.Static("/upload", "./upload")

	api := router.Group("/api")
	{
		api.POST("/sign-up/", handlers.SingUp)
		api.POST("/login/", handlers.Login)

		auth := api.Group("/auth")
		{
			account := auth.Group("/account")
			{
				account.GET("/me/", middleware.RequireAuth, handlers.Me)
			}

			// anime := auth.Group("/anime")
			// {

			// }
		}
	}

}
