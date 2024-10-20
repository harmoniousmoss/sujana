package routes

import (
	"go-scraper/handlers"

	"github.com/gofiber/fiber/v2"
)

// JobRoutes registers job-related routes.
func JobRoutes(app *fiber.App) {
	app.Get("/scrape-jobs", handlers.StoreJobs)
}
