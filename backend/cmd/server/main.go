package main

import (
	"fmt"
	"log"

	"golang/backend/internal/config"
	"golang/backend/internal/repository"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	database, err := repository.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Test the connection
	if err := repository.TestConnection(database); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}

	fmt.Println("🎉 All tests passed! Database is ready to use.")
}