package postcontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/utils"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct{}

func (pc *PostController) GetAll() ([]models.Post, error) {
	var posts []models.Post
	var fields []string = []string{"id", "username", "pic_url", "is_verified", "is_mod"}

	if err := db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pc *PostController) Get(id string) (models.Post, error) {
	var post models.Post
	var fields []string = []string{"id", "username", "pic_url", "is_verified", "is_mod"}

	if err := db.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}).First(&post, id).Error; err != nil {
		return models.Post{}, errors.New("post with given id does not exists")
	}

	return post, nil
}

func (pc *PostController) Create(c *gin.Context) (*models.Post, error) {
	user, err := controller.GetUserFromCtx(c)
	if err != nil {
		log.Println(err)
		return nil, errors.New("user not found")
	}

	isPublic, _ := strconv.ParseBool(c.PostForm("isPublic"))

	image := utils.SaveImage(c, utils.PostsImg, "image")

	post := models.Post{
		Title:    c.PostForm("title"),
		Content:  c.PostForm("content"),
		IsPublic: isPublic,
		Image:    image,
		UserID:   user.ID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		log.Println(err)
		return nil, errors.New("could not create post")
	}

	if err := db.DB.Model(&user).Association("Posts").Append(&post); err != nil {
		log.Println(err)
		return nil, errors.New("could not append to users posts")
	}

	return &post, nil
}

func (pc *PostController) Update(c *gin.Context, id string) (models.Post, error) {
	var postToUpdate models.Post

	if err := db.DB.Preload("User").First(&postToUpdate, id).Error; err != nil {
		return models.Post{}, errors.New("post with given id does not exists")
	}

	postToUpdate.Title = controller.GetOrDefault(c.PostForm("title"), postToUpdate.Title).(string)
	postToUpdate.Content = controller.GetOrDefault(c.PostForm("content"), postToUpdate.Content).(string)

	isPublic, err := strconv.ParseBool(c.PostForm("isPublic"))
	if err == nil {
		postToUpdate.IsPublic = isPublic
	}

	img, err := c.FormFile("image")
	if err == nil || img != nil {
		log.Println("Updating image")
		filepath := utils.UpdateImage(c, *postToUpdate.Image, utils.PostsImg, "image")

		postToUpdate.Image = filepath
	}

	if err := db.DB.Save(&postToUpdate).Error; err != nil {
		return models.Post{}, errors.New("error while updating post")
	}

	return postToUpdate, nil
}

func (pc *PostController) Delete(id string) error {
	if err := db.Delete(&models.Post{}, id); err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}
