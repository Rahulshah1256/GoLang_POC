package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration values
type Config struct {
	DatabaseURL       string
	SQLDriver         string
	TokenSymmetricKey string
}

// LoadConfig loads the configuration values from the environment variables
func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load .env file: %w", err)
	}

	config := Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SQLDriver:         os.Getenv("SQL_DRIVER"),
		TokenSymmetricKey: os.Getenv("TOKEN_SYMMETRIC_KEY"),
	}

	return config, nil
}
