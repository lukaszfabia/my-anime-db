package usercontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (uc *UserController) GetAll() ([]*models.User, error) {
	var users []*models.User
	var fields []string = []string{"id", "created_at", "username", "email", "pic_url"}
	if err := db.DB.Model(models.User{}).Select(fields).Order("username").Order("email").Order("id").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserController) Get(id string, props ...any) (*models.User, error) {
	var user models.User
	var fields []string = []string{"id", "created_at", "username", "email", "pic_url", "bio", "is_mod", "is_verified", "website"}

	if err := db.DB.
		Model(models.User{}).
		Preload("Friends", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email", "pic_url")
		}).
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Preload("Review", func(db *gorm.DB) *gorm.DB {
			return db.Preload("UserAnime", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Anime").Preload("User").Order("created_at DESC")
			})
		}).
		Preload("UserStats").
		Select(fields).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *UserController) Create(c *gin.Context) (*models.User, error) {
	return nil, errors.New("inserting user is not allowed")
}

func (uc *UserController) Update(c *gin.Context, id string) (*models.User, error) {
	return nil, errors.New("updating user is not allowed")
}

func (uc *UserController) Delete(id string) error {
	return errors.New("deleting user is not allowed")
}
