package main

import (
	"go-scraper/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize MongoDB
	config.InitDB()
	defer config.CloseDB() // Ensure MongoDB disconnects on exit

	// Initialize Fiber app
	app := fiber.New()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MongoDB connected!")
	})

	// Start the Fiber app on port 8080
	log.Fatal(app.Listen(":8080"))
}
