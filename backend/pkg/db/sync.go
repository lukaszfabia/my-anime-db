package db

import "api/internal/models"

func SyncDb() {
	// register models

	DB.Exec(`DO $$
		DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END $$;
	`)

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
