package db

import (
	"api/internal/models"
	"api/pkg/utils"
	"errors"
)

/*
Delete associations from object, helper function for Delete

params:

  - model: interface{} - model to delete e.g &models.Anime{}
  - assosiations: ...string - associations to delete, e.g "Genres", "Studios", "Roles"

returns:

  - error - if can't delete association
*/
func deleteAssociations(model interface{}, assosiations ...string) error {
	for _, assosiation := range assosiations {
		if err := DB.Model(model).Association(assosiation).Clear(); err != nil {
			return err
		}
	}

	return nil
}

/*
Delete object from database, provides on delete cascade, removes permanently from database

params:

  - model: interface{} - model to delete e.g &models.Anime{}
  - id: string - id of object to delete, e.g "1" usually from endpoint
  - associations: ...string - associations to delete, e.g "Genres", "Studios", "Roles"

returns:

  - error - error if something went wrong
*/
func Delete(model interface{}, id string, associations ...string) error {
	object := model

	tx := DB
	for _, assoc := range associations {
		tx = tx.Preload(assoc)
	}

	res := tx.First(object, id)
	if res.Error != nil {
		return errors.New("no model found")
	}

	if picUrlGetter, ok := object.(models.PicUrlGetter); ok {
		if picUrl := picUrlGetter.GetPicUrl(); picUrl != nil {
			utils.RemoveImage(*picUrl)
		}
	} else {
		return errors.New("cannot remove image")
	}

	if err := deleteAssociations(object, associations...); err != nil {
		return errors.New("cannot remove associations")
	}

	isDeleted := DB.Unscoped().Delete(object, id)
	if isDeleted.Error != nil {
		return errors.New("cannot remove object")
	}

	return nil
}

func Retrieve(model interface{}, wrapper interface{}, id string, associations ...string) error {
	object := wrapper

	tx := DB

	for _, association := range associations {
		tx = tx.Preload(association)
	}

	res := tx.Model(model).First(object, id)
	if res.Error != nil {
		return errors.New("no model found")
	}

	return nil
}
