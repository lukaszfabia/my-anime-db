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

		api.GET("/categories/", handlers.GetCategories)

		studio := api.Group("/studio")
		{
			studio.GET("/", handlers.GetAllStudios) // get all
			studio.GET("/:id", handlers.GetStudio)
		}

		api.GET("/all-anime/", handlers.GetAllAnimes)
		api.GET("/anime/", handlers.GetAnime)

		genre := api.Group("/genre")
		{
			genre.GET("/", handlers.GetAllGenres) // get all
			genre.GET("/:id", handlers.GetGenre)
		}

		user := api.Group("/user")
		{
			user.GET("/", handlers.GetAllUsers) // get all
			user.GET("/:id", handlers.GetUser)
		}

		actor := api.Group("/voice_actor")
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
				studio := manage.Group("/studio")
				{
					studio.POST("/", handlers.CreateStudio)
					studio.DELETE("/:id", handlers.DeleteStudio)
					studio.PUT("/:id", handlers.EditStudio)
				}

				actor := manage.Group("/voice_actor")
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
					anime.PUT("/:id", handlers.UpdateAnime)
					anime.POST("/", handlers.CreateAnime)
				}

				genre := manage.Group("/genre")
				{
					genre.DELETE("/:id", handlers.DeleteGenre)
					genre.PUT("/:id", handlers.EditGenre)
					genre.POST("/", handlers.CreateGenre)
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

			anime := auth.Group("/anime")
			{
				anime.PUT("/:id/review/", handlers.SetReview)
				anime.DELETE("/:id/remove-from-list/", handlers.DeleteFromList)
				anime.PUT("/:id/add-to-list/", handlers.AddToList)
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
