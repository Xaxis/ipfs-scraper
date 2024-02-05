package scraper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/Xaxis/ipfs-scraper/internal/db"
)

// ParseCSV parses a CSV file and returns a slice of IPFS CIDs.
func ParseCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cids := []string{}
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		cid := strings.TrimSpace(record[0])
		if validateCID(cid) {
			cids = append(cids, cid)
		}
	}
	return cids, nil
}

// validateCID checks if the provided string is a valid IPFS CID, including those with paths or extensions.
func validateCID(cid string) bool {

	// Pattern to match both CIDv0 and CIDv1, and allow for paths and extensions like ".json"
	pattern := `^(?:Qm[a-zA-Z0-9]{44}|b[a-z2-7]{58})(?:\/[\w\-\.]+)*\/?[\w\-\.]*\.json$|^Qm[a-zA-Z0-9]{44}$|^b[a-z2-7]{58}$`

	// Compile the regex pattern
	regex, _ := regexp.Compile(pattern)

	// Check if the CID matches the pattern
	return regex.MatchString(cid)
}

// ParseRawJSONToIPFSMetadata parses a raw JSON metadata string and returns the corresponding IPFSMetadata instance
func ParseRawJSONToIPFSMetadata(rawJSON string, cid string) (*db.IPFSMetadata, error) {
	var temp map[string]interface{}
	err := json.Unmarshal([]byte(rawJSON), &temp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	name, ok := temp["name"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting Name type for CID %s", cid)
	}

	description, ok := temp["description"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting Description type for CID %s", cid)
	}

	imageURL, ok := temp["image"].(string)
	if !ok {
		return nil, fmt.Errorf("error asserting ImageURL type for CID %s", cid)
	}

	// Use NewIPFSMetadata to create and return the metadata instance
	metadata := db.NewIPFSMetadata(cid, name, description, imageURL)
	return metadata, nil
}
