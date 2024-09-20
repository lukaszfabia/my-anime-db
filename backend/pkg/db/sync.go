package db

import "api/internal/models"

func SyncDb() {
	// register models
	DB.AutoMigrate(
		&models.User{},
		&models.Anime{},
		&models.Genre{},
		&models.VoiceActor{},
		&models.Character{},
		&models.Studio{},
		&models.Post{},
		&models.Role{},
		&models.OtherTitles{},
		&models.UserAnime{},
		&models.AnimeStat{},
		&models.FriendRequest{},
		&models.UserStat{},
		&models.Review{},
	)
}
