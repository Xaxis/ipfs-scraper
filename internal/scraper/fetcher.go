package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Fetcher is responsible for fetching data from IPFS.
type Fetcher struct {
	httpClient *http.Client
}

// NewFetcher creates a new Fetcher with custom http client.
func NewFetcher() *Fetcher {
	return &Fetcher{
		httpClient: &http.Client{
			Timeout: time.Second * 30, // Set a timeout for requests
		},
	}
}

// FetchMetadata makes an HTTP request to an IPFS gateway to fetch metadata for a given CID.
func (f *Fetcher) FetchMetadata(cid string) (string, error) {

	// Eventually, we will want to support fetching metadata from multiple IPFS gateways.
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)

	resp, err := f.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching CID %s: %w", cid, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code %d for CID %s", resp.StatusCode, cid)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body for CID %s: %w", cid, err)
	}

	return string(body), nil
}
