package genrecontroller

import (
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/validators"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GenreController struct{}

func (gc *GenreController) GetAll() ([]*models.Genre, error) {
	var genres []*models.Genre

	if db.DB.Find(&genres).Error != nil {
		return nil, errors.New("can not find model")
	}

	return genres, nil
}

func (gc *GenreController) Get(id string, props ...any) (*models.Genre, error) {
	var genre models.Genre

	if db.DB.First(&genre, id).Error != nil {
		return nil, errors.New("can not find model")
	}

	return &genre, nil
}

func (gc *GenreController) Delete(id string) error {
	if err := db.DB.Delete(&models.Genre{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (gc *GenreController) Create(c *gin.Context) (*models.Genre, error) {
	genre := c.PostForm("genre")
	if validators.IsEmpty(genre) {
		return nil, errors.New("genre is required")
	}

	var tmp models.Genre

	// restore genre if it was deleted
	if db.DB.Unscoped().First(&tmp, "name = ?", genre).Error == nil && tmp.DeletedAt.Valid {
		tmp.DeletedAt = gorm.DeletedAt{}
		if db.DB.Save(&tmp).Error != nil {
			return nil, errors.New("genre already exists")
		}

		return &tmp, nil
	}

	var newGenre models.Genre = models.Genre{
		Name: genre,
	}

	if db.DB.Save(&newGenre).Error != nil {
		return nil, errors.New("can not create model")
	}

	return &newGenre, nil
}

func (gc *GenreController) Update(c *gin.Context, id string) (*models.Genre, error) {
	return nil, errors.New("not implemented")
}
