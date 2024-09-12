package voiceactorcontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/pkg/db"
	"api/pkg/utils"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

type VoiceActorController struct {
}

func (ac *VoiceActorController) GetAll() ([]*models.VoiceActor, error) {
	var voiceActors []*models.VoiceActor

	if err := db.DB.Unscoped().
		Order("lastname, name, id").
		Find(&voiceActors).Error; err != nil {
		return nil, errors.New("failed to retrieve voice actors")
	}
	return voiceActors, nil
}

func (ac *VoiceActorController) Get(id string, props ...any) (*models.VoiceActor, error) {
	var voiceActor models.VoiceActor

	if err := db.DB.
		Preload("Roles").
		First(&voiceActor, id).Error; err != nil {
		return nil, errors.New("failed to retrieve voice actor")
	}
	return &voiceActor, nil
}

func (ac *VoiceActorController) Create(c *gin.Context) (*models.VoiceActor, error) {

	picUrl := utils.SaveImage(c, utils.ActorsImg, "pic")

	var voiceActor models.VoiceActor = models.VoiceActor{
		Name:      c.PostForm("name"),
		LastName:  c.PostForm("lastname"),
		Birthdate: c.PostForm("birthdate"),
		PicUrl:    picUrl,
	}

	if err := db.DB.Create(&voiceActor).Error; err != nil {
		log.Println(err)
		return nil, errors.New("failed to create voice actor")
	}

	return &voiceActor, nil
}

func (ac *VoiceActorController) Update(c *gin.Context, id string) (*models.VoiceActor, error) {
	var actorToUpdate models.VoiceActor

	if err := db.DB.First(&actorToUpdate, id).Error; err != nil {
		return nil, errors.New("failed to find voice actor")
	}

	actorToUpdate.Name = controller.GetOrDefault(c.PostForm("name"), actorToUpdate.Name).(string)
	actorToUpdate.LastName = controller.GetOrDefault(c.PostForm("lastname"), actorToUpdate.LastName).(string)
	actorToUpdate.Birthdate = controller.GetOrDefault(c.PostForm("birthdate"), actorToUpdate.Birthdate).(string)
	actorToUpdate.PicUrl = utils.UpdateImage(c, actorToUpdate.PicUrl, utils.ActorsImg, "pic")

	if db.DB.Save(&actorToUpdate).Error != nil {
		return nil, errors.New("failed to update voice actor")
	}

	return &actorToUpdate, nil
}

func (ac *VoiceActorController) Delete(id string) error {
	if err := db.Delete(&models.VoiceActor{}, id, db.Association{
		Model: "Roles",
	}); err != nil {
		return err
	}

	return nil
}
