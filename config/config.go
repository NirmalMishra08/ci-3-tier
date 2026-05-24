package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func Load() (*Config, error) {
	godotenv.Load("./.env")
	godotenv.Load("../.env")

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	return &Config{
		PORT: port,
	}, nil
}
