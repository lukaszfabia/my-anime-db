package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	CorsConfig cors.Config
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	const sep string = ","

	corsConfig := cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), sep), // Dostosuj do swoich potrzeb
		AllowMethods:     strings.Split(os.Getenv("ALLOW_METHODS"), sep),
		AllowHeaders:     strings.Split(os.Getenv("ALLOW_HEADERS"), sep),
		ExposeHeaders:    strings.Split(os.Getenv("EXPOSE_HEADERS"), sep),
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return &Config{
		Port:       os.Getenv("API_PORT"),
		CorsConfig: corsConfig,
	}
}
