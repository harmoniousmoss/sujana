package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBClient holds the connected MongoDB client.
var DBClient *mongo.Database

// InitDB initializes the MongoDB connection and stores the client globally.
func InitDB() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve MongoDB URI from environment variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set in .env")
	}

	// Connect to MongoDB directly using mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Verify the connection with Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")
	DBClient = client.Database("jobsdb") // Store the database client
}

// CloseDB ensures the MongoDB connection is properly closed.
func CloseDB() {
	if DBClient == nil {
		return // No connection to close
	}
	if err := DBClient.Client().Disconnect(context.Background()); err != nil {
		log.Fatal("Error disconnecting MongoDB:", err)
	}
	log.Println("Disconnected from MongoDB.")
}
