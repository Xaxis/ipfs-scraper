package db

type IPFSMetadata struct {
	CID         string `json:"cid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
}

func NewIPFSMetadata(cid, name, description, imageURL string) *IPFSMetadata {
	return &IPFSMetadata{
		CID:         cid,
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
	}
}
