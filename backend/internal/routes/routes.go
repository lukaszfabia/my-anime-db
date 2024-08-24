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
			user.GET("/:id", handlers.GetUser)
		}

		actor := api.Group("/voice-actor")
		{
			actor.GET("/", handlers.GetAllVoiceActors) // get all
			actor.GET("/:id", handlers.GetVoiceActor)
		}

		character := api.Group("/character")
		{
			character.GET("/", handlers.GetAllCharacters) // get all
			character.GET("/:id", handlers.GetCharacter)
		}

		post := api.Group("/post")
		{
			post.GET("/", handlers.GetAllPosts) // get all
			post.GET("/:id", handlers.GetPost)
		}

		search := api.Group("/search")
		{
			search.GET("/anime/:query")
			search.GET("/character/:query")
			search.GET("/user/:query")
		}

		auth := api.Group("/auth", middleware.RequireAuth)
		{
			manage := auth.Group("/manage", middleware.ReqiureMod)
			{
				actor := manage.Group("/voice-actor")
				{
					actor.POST("/", handlers.CreateVoiceActor)
					actor.DELETE("/:id", handlers.DeleteVoiceActor)
					actor.PUT("/:id", handlers.EditVoiceActor)
				}

				character := manage.Group("/character")
				{
					character.DELETE("/:id", handlers.DeleteCharacter)
					character.PUT("/:id", handlers.EditCharacter)
					character.POST("/", handlers.CreateCharacter)
				}

				anime := manage.Group("/anime")
				{
					anime.DELETE("/:id", handlers.DeleteAnime)
					anime.PUT("/:id", handlers.EditAnime)
					anime.POST("/", handlers.CreateAnime)
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
				friend.DELETE("/:id", handlers.DeleteFriend)

				friend.GET("/invitations/", handlers.GetInvitations)

				friend.GET("/state/", handlers.GetFriendState)
				// merge it
				friend.POST("/:id/respond/", handlers.RespondToFriendRequest)
			}

		}

	}

}
