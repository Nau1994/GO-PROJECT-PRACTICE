package config

import (
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	Port        string
	DatabaseURL string
}

// LoadConfig loads configuration from environment variables and returns an error if something is missing
func LoadConfig() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is not set")
	}

	return Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}, nil
}
