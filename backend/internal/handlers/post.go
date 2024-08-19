package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/pkg/db"
	"api/pkg/utils"
	"api/pkg/validators"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrievePost(c *gin.Context) {
	var id string = c.Param("id")
	var post models.Post
	userSelector := db.ToSelectFunc(db.DB, "id", "username", "pic_url", "is_verified", "is_mod")
	if err := db.Retrieve(models.Post{}, &post, id, db.Association{Model: "User", Selector: userSelector}); err != nil {
		c.JSON(http.StatusNotFound, post)
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	userObj, _ := c.Get("user")
	user, _ := userObj.(models.User)

	var postForm models.PostForm
	if !validators.IsFormDataValid(c, &postForm) {
		c.JSON(http.StatusBadRequest, response.BadForm())
		return
	}

	isPublic, err := strconv.ParseBool(c.PostForm("isPublic"))

	if err != nil {
		log.Println(err)
	}

	log.Println(isPublic)

	var imageFilename string
	if file, err := c.FormFile("image"); file != nil || err == nil {
		imageFilename = utils.SaveImage(c, utils.SaverProps{
			Dir:         utils.PostsImg,
			Placeholder: utils.DefaultConent,
			KeyToImg:    "image",
			Filename:    fmt.Sprintf("%s-%s", c.PostForm("title"), user.Username),
		})
	}

	post := models.Post{
		Title:    c.PostForm("title"),
		Content:  c.PostForm("content"),
		IsPublic: isPublic,
		Image:    &imageFilename,
		UserID:   user.ID,
	}

	log.Println(post)

	if err := db.DB.Create(&post).Error; err != nil {
		msgErr := "could not create post"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	if err := db.DB.Model(&user).Association("Posts").Append(&post); err != nil {
		msgErr := "could not append to users posts"
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusCreated, post)
}

func DeletePost(c *gin.Context) {
	var id string = c.Param("id")

	if err := db.Delete(&models.Post{}, id); err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	msg := "post has been removed"
	c.JSON(http.StatusOK, response.NewResponse(&msg, nil))
}

func EditPost(c *gin.Context) {
	var id string = c.Param("id")

	var postToEdit models.Post

	if err := db.DB.Preload("User").First(&postToEdit, id).Error; err != nil {
		msgErr := "post with given id does not exists"
		c.JSON(http.StatusNotFound, response.NewResponse(nil, &msgErr))
		return
	}

	postToEdit.Title = c.DefaultPostForm("title", postToEdit.Title)
	postToEdit.Content = c.DefaultPostForm("content", postToEdit.Content)
	isPublic, err := strconv.ParseBool(c.PostForm("isPublic"))
	if err == nil {
		postToEdit.IsPublic = isPublic
	}

	img, err := c.FormFile("image")
	if err == nil || img != nil {
		utils.RemoveImage(*postToEdit.Image)

		filepath := utils.SaveImage(c, utils.SaverProps{
			Dir:         utils.PostsImg,
			Placeholder: utils.DefaultConent,
			KeyToImg:    "image",
			Filename:    fmt.Sprintf("%s-%s", postToEdit.User.Username, c.DefaultPostForm("title", postToEdit.Title)),
		})

		postToEdit.Image = &filepath
	}

	db.DB.Save(&postToEdit)

	c.JSON(http.StatusOK, postToEdit)
}

func GetAllPosts(c *gin.Context) {
	var allPosts []models.Post
	userSelector := db.ToSelectFunc(db.DB, "id", "username", "pic_url", "is_mod", "is_verified")
	if err := db.RetrieveAll(&models.Post{}, &allPosts, db.ToOrder(db.DB, "created_at DESC"), db.Association{Model: "User", Selector: userSelector}); err != nil {
		msgErr := err.Error()
		c.JSON(http.StatusBadRequest, response.NewResponse(nil, &msgErr))
		return
	}

	c.JSON(http.StatusOK, allPosts)
}
