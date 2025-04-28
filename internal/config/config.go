package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBURL         string
	Environment   string
	AuthSecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		DBURL:         os.Getenv("DATABASE_URL"),
		Port:          os.Getenv("PORT"),
		Environment:   os.Getenv("ENVIRONMENT"),
		AuthSecretKey: os.Getenv("AUTH_SECRET_KEY"),
	}

	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return cfg
}
