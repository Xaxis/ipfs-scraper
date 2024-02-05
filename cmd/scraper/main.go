package main

import (
	"fmt"
	"log"
	"os"
	"time"

	db "github.com/Xaxis/ipfs-scraper/internal/db"
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

	var dbInstance *db.Database
	var err error

	// Attempt to connect to the db with retries
	for attempt := 1; attempt <= 5; attempt++ {
		dbInstance, err = db.NewDatabase(dsn)
		if err != nil {
			log.Printf("Failed to connect to db: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	// Check if the db connection was successful
	if err != nil {
		log.Fatalf("Error connecting to db after retries: %v", err)
	}

	defer dbInstance.Close()

	// Run scraper instance using CIDs from the csv file
	ipfsScraper := sc.NewFetcher()
	csvFile := "./data/ipfs_cids.csv"
	s := sc.NewScraper(ipfsScraper, dbInstance, csvFile)
	s.Scrape()
}
