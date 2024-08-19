package db

import (
	"api/internal/models"
	"api/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type Association struct {
	Model    string
	Selector func(db *gorm.DB) *gorm.DB
}

func ToSelectFunc(db *gorm.DB, values ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(values)
	}
}

func ToOrder(db *gorm.DB, field string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(field)
	}
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
		if err := DB.Model(model).Association(assosiation.Model).Clear(); err != nil {
			return err
		}
	}

	return nil
}

/*
Delete strong entity from database, provides on delete cascade, removes permanently from database

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

func loadAssociations(associations ...Association) *gorm.DB {
	tx := DB
	for _, association := range associations {
		tx = tx.Preload(association.Model, func(db *gorm.DB) *gorm.DB {
			if association.Selector != nil {
				return association.Selector(db)
			}
			return db
		})
	}

	return tx
}

func Retrieve(model interface{}, dest interface{}, id string, associations ...Association) error {
	object := dest

	tx := loadAssociations(associations...)

	res := tx.Model(model).First(object, id)
	if res.Error != nil {
		return errors.New("no model found")
	}

	return nil
}

func RetrieveAll(model interface{}, dest interface{}, customOrder func(db *gorm.DB) *gorm.DB, associations ...Association) error {

	tx := loadAssociations(associations...)

	res := tx.Model(model).Order(customOrder(tx)).Find(dest)

	if res.Error != nil {
		return errors.New("no model found")
	}

	return nil
}
