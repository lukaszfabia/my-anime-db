package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	TrustedProxies []string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	proxies := strings.Split(os.Getenv("TRUSTED_PROXIES"), ",") // idk why it doesnt work

	return &Config{
		Port:           os.Getenv("API_PORT"),
		TrustedProxies: proxies,
	}
}
