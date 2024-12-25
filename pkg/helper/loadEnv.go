package helper

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func LoadEnv() {
	// Load environment variables from the .env file
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
