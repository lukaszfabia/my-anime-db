package main

import (
	"api/internal/config"
	"api/internal/server"
	"api/pkg/db"
	"log"
)

func init() {
	db.ConnectToDb()
	db.SyncDb()
}

func main() {
	cfg := config.Load()
	s := server.New(cfg)
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
