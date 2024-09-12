package db

import (
	"api/internal/models"
	"api/pkg/utils"
	"errors"
	"log"
)

type Association struct {
	Model string
}

/*
Delete associations from object, helper function for Delete

params:

  - model: interface{} - model to delete e.g &models.Anime{}
  - associations: ...string - associations to delete, e.g "Genres", "Studios", "Roles"

returns:

  - error - if can't delete association
*/
func deleteAssociations(model interface{}, assosiations ...Association) error {
	for _, assosiation := range assosiations {
		if err := DB.Unscoped().Model(model).Association(assosiation.Model).Clear(); err != nil {
			return err
		}
	}

	return nil
}

/*
Delete entity from database, provides on delete cascade, removes from database

params:

  - model: interface{} - model to delete e.g &models.Anime{}
  - id: string - id of object to delete, e.g "1" usually from endpoint
  - associations: ...string - associations to delete, e.g "Genres", "Studios", "Roles"

returns:

  - error - error if something went wrong
*/
func Delete(model interface{}, id string, associations ...Association) error {
	object := model

	tx := DB

	for _, assoc := range associations {
		tx = tx.Preload(assoc.Model)
	}

	res := tx.First(object, id)
	if res.Error != nil {
		return errors.New("no model found")
	}

	if err := deleteAssociations(object, associations...); err != nil {
		log.Println(err.Error())
		return errors.New("cannot remove associations")
	}

	isDeleted := DB.Unscoped().Delete(object, id)
	if isDeleted.Error != nil {
		log.Println(isDeleted.Error.Error())
		return errors.New("cannot remove object")
	}

	if picUrlGetter, ok := object.(models.PicUrlGetter); ok {
		if picUrl := picUrlGetter.GetPicUrl(); picUrl != nil {
			utils.RemoveImage(picUrl)
		}
	}

	return nil
}
