package main

import (
	"log"

	"github.com/felipe-lima-coelho/desafio-taghos-backend-jr/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db := config.Database()
}
