package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	DBURL       string
	Environment string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBURL:       os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
	}

	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return cfg
}
