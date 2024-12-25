package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
}

// getEnv retrieves the value of the environment variable named by the key.
func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

// getEnvInt retrieves the value of the environment variable named by the key and converts it to an integer.
func getEnvInt(key string) int {
	value := getEnv(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error converting environment variable %s to int: %v", key, err)
	}
	return intValue
}

// LoadConfig loads configuration from a .env file located in the configs directory.
func LoadConfig() *Config {
	envFilePath := filepath.Join("configs", ".env")

	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBHost: getEnv("DB_HOST"),
		DBPort: getEnvInt("DB_PORT"),
		DBUser: getEnv("DB_USER"),
		DBPass: getEnv("DB_PASSWORD"),
		DBName: getEnv("DB_NAME"),
	}
}
