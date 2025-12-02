package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}