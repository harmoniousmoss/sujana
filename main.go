// main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Create a route for "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go Scraper")
	})

	// Run the Fiber app with default settings (port 8080)
	log.Fatal(app.Listen(""))
}
