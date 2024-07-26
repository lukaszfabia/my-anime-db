package db

import "api/internal/models"

func SyncDb() {
	// register models here ...
	DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Studio{},
		&models.Genre{},
		&models.Anime{},
		&models.VoiceActor{},
		&models.Character{},
		&models.UserAnime{},
		&models.Roles{},
	)
}
