package db

import "api/internal/models"

func setConsts() {
	// options := []string{
	// 	"action",
	// 	"cyberpunk",
	// 	"drama",
	// 	"ecchi",
	// 	"experimental",
	// 	"fantasy",
	// 	"harem",
	// 	"hentai",
	// 	"historical",
	// 	"horror",
	// 	"comedy",
	// 	"crime",
	// 	"magic",
	// 	"mecha",
	// 	"male-harem",
	// 	"music",
	// 	"supernatural",
	// 	"madness",
	// 	"slice-of-life",
	// 	"parody",
	// 	"adventure",
	// 	"psychological",
	// 	"romance",
	// 	"romance-separated",
	// 	"sci-fi",
	// 	"shoujo-ai",
	// 	"shounen-ai",
	// 	"space-opera",
	// 	"sports",
	// 	"steampunk",
	// 	"school",
	// 	"martial-arts",
	// 	"mystery",
	// 	"thriller",
	// 	"military",
	// 	"yaoi",
	// 	"yuri",
	// }

	// for _, opt := range options {
	// 	DB.Create(&models.Genre{Name: models.GenreOption(opt)})
	// }

	DB.Exec("INSERT INTO studios (name, website) VALUES  ('Madhouse', 'https://www.madhouse.co.jp/index.html'), ('Studio Bones', 'https://www.bones.co.jp/'), ('Wit Studio', 'https://www.witstudio.co.jp/'), ('Toei Animation', 'https://corp.toei-anim.co.jp/en/index.html'), ('MAPPA Studio', 'https://www.mappa.co.jp/en/'), ('Studio Ghibli', ''), ('Sunrise Studio', 'https://www.sunrise-inc.co.jp/international/'), ('A-1 Pictures', 'https://a1p.jp/'), ('Ufotable', 'https://www.ufotable.com/en/'), ('Studio Pierrot','https://en.pierrot.jp/'),('Production I.G', 'https://www.productionig.com/'), ('Studio Trigger', 'https://www.st-trigger.co.jp/'), ('P.A.Works', 'https://www.pa-works.jp/en/'), ('CoMix Wave','https://www.cwfilms.jp/en/')")
}

func SyncDb() {
	// register models

	// setConsts()

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
	)
}
