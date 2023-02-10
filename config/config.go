package config

import (
	"os"
)

// Config holds the configuration variables for the application
type Config struct {
	PostgresDBUser     string
	PostgresDBPassword string
	PostgresDBHost     string
	PostgresDBName     string
}

// GetConfig returns the configuration variables for the application
func GetConfig() *Config {
	return &Config{
		PostgresDBUser:     getEnv("POSTGRES_DB_USER", ""),
		PostgresDBPassword: getEnv("POSTGRES_DB_PASSWORD", ""),
		PostgresDBHost:     getEnv("POSTGRES_DB_HOST", ""),
		PostgresDBName:     getEnv("POSTGRES_DB_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
