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

		anime := api.Group("/anime")
		{
			anime.GET("/") // get all
			anime.GET("/:id")
		}
		user := api.Group("/user")
		{
			user.GET("/", handlers.GetAllUsers) // get all
			user.GET("/:id", handlers.RetriveUser)
		}

		actors := api.Group("/actors")
		{
			actors.GET("/") // get all
			actors.GET("/:id")
		}

		character := api.Group("/character")
		{
			character.GET("/") // get all
			character.GET("/:id")
		}

		search := api.Group("/search")
		{
			search.GET("/anime/:query")
			search.GET("/character/:query")
			search.GET("/user/:query")
		}

		auth := api.Group("/auth", middleware.RequireAuth)
		{
			account := auth.Group("/account")
			{
				account.GET("/me/", handlers.Me)
				account.DELETE("/me/", handlers.DeleteMe)
				account.PUT("/me/", handlers.EditMe)
			}

			review := auth.Group("/review")
			{
				review.DELETE("/:id")
				review.PUT("/:id")
				review.POST("/")
			}

			post := auth.Group("/post")
			{
				post.DELETE("/:id")
				post.PUT("/:id")
				post.POST("/")
			}

			anime := auth.Group("/anime")
			{
				anime.DELETE("/:id")
				anime.PUT("/:id")
				anime.POST("/")
			}

			actors := auth.Group("/actor")
			{
				actors.DELETE("/:id")
				actors.PUT("/:id")
				actors.POST("/")
			}

			characters := auth.Group("/characters")
			{
				characters.DELETE("/:id")
				characters.PUT("/:id")
				characters.POST("/")
			}

		}

	}

}
