package main

import (
	"fmt"
	"log"
	"os"
	"time"

	sc "github.com/Xaxis/ipfs-scraper/internal/scraper"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	var db *sc.Database
	var err error

	// Attempt to connect to the database with retries
	for attempt := 1; attempt <= 5; attempt++ {
		db, err = sc.NewDatabase(dsn)
		if err != nil {
			log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	// Check if the database connection was successful
	if err != nil {
		log.Fatalf("Error connecting to database after retries: %v", err)
	}

	defer db.Close()

	// Proceed with the rest of your application logic
	ipfsScraper := sc.NewFetcher()
	csvFile := "./data/ipfs_cids.csv"

	s := sc.NewScraper(ipfsScraper, db, csvFile)
	s.Scrape()
}
