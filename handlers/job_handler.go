package handlers

import (
	"context"
	"log"
	"time"

	"go-scraper/config"
	scraper "go-scraper/scrapper"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// StoreJobs scrapes jobs and stores them in MongoDB.
func StoreJobs(c *fiber.Ctx) error {
	// Scrape jobs from the target site
	jobs, err := scraper.ScrapeJobs()
	if err != nil {
		return c.Status(500).SendString("Failed to scrape jobs: " + err.Error())
	}

	// Get MongoDB collection reference
	db := config.DBClient
	collection := db.Collection("remote_jobs")

	// Insert jobs into MongoDB
	for _, job := range jobs {
		_, err := collection.InsertOne(context.Background(), bson.M{
			"title": job.Title,
			"link":  job.Link,
			"date":  time.Now(),
		})
		if err != nil {
			log.Println("Failed to insert job:", err)
			continue
		}
	}

	return c.SendString("Jobs stored successfully!")
}
