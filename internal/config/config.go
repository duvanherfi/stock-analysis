package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL      string
	APIKey      string
	DatabaseURL string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		APIURL:      os.Getenv("API_URL"),
		APIKey:      os.Getenv("API_KEY"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
