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
			anime.GET("/", handlers.GetAllAnimes) // get all
			anime.GET("/:id", handlers.RetrieveAnime)
		}
		user := api.Group("/user")
		{
			user.GET("/", handlers.GetAllUsers) // get all
			user.GET("/:id", handlers.RetrieveUser)
		}

		actors := api.Group("/actors")
		{
			actors.GET("/", handlers.GetAllActors) // get all
			actors.GET("/:id", handlers.RetrieveActor)
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
				anime.DELETE("/:id", handlers.RemoveAnime)
				anime.PUT("/:id", handlers.SaveOrUpdateAnime)
				anime.POST("/", handlers.SaveOrUpdateAnime)
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

			friend := auth.Group("/friend")
			{
				friend.POST("/:id", handlers.AddFriend)
				friend.DELETE("/:id", handlers.RemoveFriend)

				friend.POST("/:id/respond/", handlers.RespondToFriendRequest)

				friend.GET("/invitations/", handlers.GetInvitations)
			}

		}

	}

}
