package handlers

import (
	"api/internal/app"
	"api/internal/controller"
	postcontroller "api/internal/controller/post_controller"
	"api/internal/models"
	"api/pkg/validators"
	postvalidator "api/pkg/validators/post_validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var postController controller.Controller[models.Post] = &postcontroller.PostController{}
var pv validators.Validator = &postvalidator.PostValidator{}

func GetPost(c *gin.Context) {
	r := app.Gin{Ctx: c}

	post, err := postController.Get(c.Param("id"))

	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, post)
}

func CreatePost(c *gin.Context) {
	r := app.Gin{Ctx: c}

	if !pv.Validate(c) {
		log.Println("Validation failed")
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if newPost, err := postController.Create(c); err != nil {
		log.Println(err)
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusCreated, app.Ok, newPost)
	}
}

func DeletePost(c *gin.Context) {
	r := app.Gin{Ctx: c}
	id := c.Param("id")

	if err := postController.Delete(id); err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, nil)
}

func EditPost(c *gin.Context) {
	var id string = c.Param("id")

	r := app.Gin{Ctx: c}

	if !pv.Validate(c) {
		log.Println("Validation failed")
		r.NewResponse(http.StatusBadRequest, app.InvalidData, nil)
		return
	}

	if editedPost, err := postController.Update(c, id); err != nil {
		log.Println(err)
		r.NewResponse(http.StatusInternalServerError, app.Failed, nil)
		return
	} else {
		r.NewResponse(http.StatusOK, app.Ok, editedPost)
	}

}

func GetAllPosts(c *gin.Context) {
	r := app.Gin{Ctx: c}

	posts, err := postController.GetAll()

	if err != nil {
		r.NewResponse(http.StatusNotFound, app.Failed, nil)
		return
	}

	r.NewResponse(http.StatusOK, app.Ok, posts)
}
