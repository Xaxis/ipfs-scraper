package scraper

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

// NewDatabase establishes a new database connection
func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{DB: db}, nil
}

// SaveMetaData saves an IPFSMetadata instance into the database
func (db *Database) SaveMetaData(metadata *IPFSMetadata) error {
	query := `INSERT INTO metadata (CID, Name, Description, ImageURL) VALUES ($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(metadata.CID, metadata.Name, metadata.Description, metadata.ImageURL)
	return err
}

// Modified version of SaveMetaData to prevent duplicate entries
//func (db *Database) SaveMetaData(metadata *IPFSMetadata) error {
//	query := `INSERT INTO metadata (CID, Name, Description, ImageURL) VALUES ($1, $2, $3, $4) ON CONFLICT (CID) DO NOTHING`
//	_, err := db.Exec(query, metadata.CID, metadata.Name, metadata.Description, metadata.ImageURL)
//	return err
//}
