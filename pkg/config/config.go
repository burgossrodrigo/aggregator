package config

import (
	"log"
	"os"

	models "aggregator/pkg/models"

	"github.com/joho/godotenv"
)

func LoadEnv() (cfg models.Configs) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg.ZeroXAPIKey = os.Getenv("ZEROXAPIKEY")

	return cfg
}
