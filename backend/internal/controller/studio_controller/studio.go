package studiocontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type StudioController struct {
}

func (s *StudioController) GetAll() ([]*models.Studio, error) {
	var studios []*models.Studio

	if err := db.DB.Preload("Anime").Find(&studios).Error; err != nil {
		return nil, err
	}

	return studios, nil
}

func (s *StudioController) Get(id string, props ...any) (*models.Studio, error) {
	var studio models.Studio

	if err := db.DB.Preload("Anime").First(&studio, id).Error; err != nil {
		return nil, err
	}

	return &studio, nil
}

func (s *StudioController) Create(c *gin.Context) (*models.Studio, error) {
	date, err := time.Parse(time.DateOnly, c.PostForm("establishedDate"))
	if err != nil {
		date = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	logo := utils.SaveImage(c, "logo", "studio")

	var newStudio models.Studio = models.Studio{
		Name:            c.PostForm("name"),
		EstablishedDate: &date,
		Website:         controller.GetOrDefault(c.PostForm("website"), "").(string),
		LogoUrl:         logo,
	}

	if err := db.DB.Create(&newStudio).Error; err != nil {
		return nil, err
	}

	return &newStudio, nil
}

func (s *StudioController) Update(c *gin.Context, id string) (*models.Studio, error) {
	var studioToUpdate models.Studio

	if err := db.DB.First(&studioToUpdate, id).Error; err != nil {
		return nil, err
	}

	date, err := time.Parse(time.DateOnly, c.PostForm("establishedDate"))

	if err == nil {
		studioToUpdate.EstablishedDate = &date
	}

	pic := utils.SaveImage(c, utils.StudiosLogo, "pic")

	if pic != nil {
		studioToUpdate.LogoUrl = pic
	}

	studioToUpdate.Name = controller.GetOrDefault(c.PostForm("name"), studioToUpdate.Name).(string)
	studioToUpdate.Website = controller.GetOrDefault(c.PostForm("website"), studioToUpdate.Website).(string)

	if db.DB.Save(&studioToUpdate).Error != nil {
		return nil, err
	}

	return &studioToUpdate, nil
}

func (s *StudioController) Delete(id string) error {
	return db.Delete(&models.Studio{}, id, db.Association{Model: "Anime"})
}
