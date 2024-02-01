package scraper

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
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
func ParseRawJSONToIPFSMetadata(rawJSON string) (*IPFSMetadata, error) {
	metadata := &IPFSMetadata{}
	err := json.Unmarshal([]byte(rawJSON), metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
