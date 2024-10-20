package main

import (
	"log"
	"os"

	"go-scraper/config"
	"go-scraper/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the MongoDB connection
	config.InitDB()
	defer config.CloseDB() // Ensure MongoDB disconnects on exit

	// Initialize Fiber app
	app := fiber.New()

	// Register routes
	routes.JobRoutes(app)

	// Start the Fiber app on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
