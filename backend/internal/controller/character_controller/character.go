package charactercontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type CharacterController struct {
}

func (cc *CharacterController) GetAll() ([]*models.Character, error) {
	var characters []*models.Character

	if err := db.DB.
		Order("last_name, name, id").
		Find(&characters).Error; err != nil {
		return nil, errors.New("failed to retrieve voice actors")
	}
	return characters, nil
}

func (cc *CharacterController) Get(id string, props ...any) (*models.Character, error) {
	var character models.Character

	if err := db.DB.
		Preload("Roles").
		First(&character, id).Error; err != nil {
		return nil, errors.New("failed to retrieve voice actor")
	}
	return &character, nil
}

func (cc *CharacterController) Create(c *gin.Context) (*models.Character, error) {

	picUrl := utils.SaveImage(c, utils.CharactersImg, "pic")

	var character models.Character = models.Character{
		Name:        c.PostForm("name"),
		LastName:    c.PostForm("lastname"),
		Information: c.PostForm("information"),
		PicUrl:      picUrl,
	}

	if err := db.DB.Create(&character).Error; err != nil {
		return nil, errors.New("failed to create voice actor")
	}

	return &character, nil
}

func (cc *CharacterController) Update(c *gin.Context, id string) (*models.Character, error) {
	var characterToUpdate models.Character

	if err := db.DB.First(&characterToUpdate, id).Error; err != nil {
		return nil, errors.New("failed to find voice actor")
	}

	characterToUpdate.Name = controller.GetOrDefault(c.PostForm("name"), characterToUpdate.Name).(string)
	characterToUpdate.LastName = controller.GetOrDefault(c.PostForm("lastname"), characterToUpdate.LastName).(string)
	characterToUpdate.Information = controller.GetOrDefault(c.PostForm("information"), characterToUpdate.Information).(string)
	characterToUpdate.PicUrl = utils.UpdateImage(c, characterToUpdate.PicUrl, utils.CharactersImg, "pic")

	if db.DB.Save(&characterToUpdate).Error != nil {
		return nil, errors.New("failed to update voice actor")
	}

	return &characterToUpdate, nil
}

func (cc *CharacterController) Delete(id string) error {
	if err := db.Delete(&models.Character{}, id, db.Association{
		Model: "Roles",
	}); err != nil {
		return err
	}

	return nil
}
