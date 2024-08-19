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

		actor := api.Group("/actors")
		{
			actor.GET("/", handlers.GetAllActors) // get all
			actor.GET("/:id", handlers.RetrieveActor)
		}

		character := api.Group("/character")
		{
			character.GET("/") // get all
			character.GET("/:id")
		}

		post := api.Group("/post")
		{
			post.GET("/", handlers.GetAllPosts) // get all
			post.GET("/:id", handlers.RetrievePost)
		}

		search := api.Group("/search")
		{
			search.GET("/anime/:query")
			search.GET("/character/:query")
			search.GET("/user/:query")
		}

		auth := api.Group("/auth", middleware.RequireAuth)
		{
			edit := auth.Group("/edit", middleware.ReqiureMod)
			{
				anime := edit.Group("/anime")
				{
					anime.DELETE("/:id", handlers.RemoveAnime)
					anime.PUT("/:id", handlers.SaveOrUpdateAnime)
					anime.POST("/", handlers.SaveOrUpdateAnime)
				}

				actor := edit.Group("/actor")
				{
					actor.DELETE("/:id")
					actor.PUT("/:id")
					actor.POST("/")
				}

				character := edit.Group("/characters")
				{
					character.DELETE("/:id")
					character.PUT("/:id")
					character.POST("/")
				}
			}

			account := auth.Group("/account")
			{
				account.GET("/me/", handlers.Me)
				account.DELETE("/me/", handlers.DeleteMe)
				account.PUT("/me/", handlers.EditMe)
				account.POST("/send-code/", middleware.ForNotVerified, handlers.SendCode)
				account.POST("/verify/", middleware.ForNotVerified, handlers.Verify)
			}

			review := auth.Group("/review")
			{
				review.DELETE("/:id")
				review.PUT("/:id")
				review.POST("/")
			}

			post := auth.Group("/post")
			{
				post.DELETE("/:id", handlers.DeletePost)
				post.PUT("/:id", handlers.EditPost)
				post.POST("/", handlers.CreatePost)
			}

			friend := auth.Group("/friend")
			{
				friend.POST("/:id", handlers.AddFriend)
				friend.DELETE("/:id", handlers.RemoveFriend)

				friend.GET("/invitations/", handlers.GetInvitations)

				friend.GET("/state/", handlers.GetFriendState)
				// merge it
				friend.POST("/:id/respond/", handlers.RespondToFriendRequest)
			}

		}

	}

}
