package main

import (
	"fmt"
	"github.com/Xaxis/ipfs-scraper/internal/db"
	"github.com/Xaxis/ipfs-scraper/internal/server"
	"log"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Construct DSN from environment variables
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	// Initialize the db. SHOULDN'T need to retry because our start.sh order, in combination with scraper retries,
	//should ensure the db is up.
	db, err := db.NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize the db: %v", err)
	}

	// Create a new server instance
	apiServer := server.NewServer(db)

	// Start the API server
	apiServer.Start(":8080")
}
