package scraper

import (
	"log"
)

type Scraper struct {
	ipfsScraper *Fetcher
	csvFile     string
	db          *Database
}

func NewScraper(ipfsScraper *Fetcher, db *Database, csvFile string) *Scraper {
	return &Scraper{
		ipfsScraper: ipfsScraper,
		db:          db,
		csvFile:     csvFile,
	}
}

func (s *Scraper) Scrape() {
	cids, err := ParseCSV(s.csvFile)
	if err != nil {
		log.Fatalf("Failed to parse CSV file: %s", err)
		return
	}

	for _, cid := range cids {
		rawData, err := s.ipfsScraper.FetchMetadata(cid)
		if err != nil {
			log.Printf("Failed to fetch metadata for CID %s: %s", cid, err)
			continue
		}

		metadata, err := ParseRawJSONToIPFSMetadata(rawData, cid)
		if err != nil {
			log.Printf("Failed to parse metadata for CID %s: %s", cid, err)
			continue
		}

		err = s.db.SaveMetaData(metadata)
		if err != nil {
			log.Printf("Failed to save metadata for CID %s: %s", cid, err)
		} else {
			log.Printf("Successfully saved metadata for CID %s", cid)
		}
	}
}
