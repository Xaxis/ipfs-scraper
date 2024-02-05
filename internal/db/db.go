package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

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

// SaveMetaData saves an IPFSMetadata instance into the db
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

// GetAllTokens retrieves all tokens from the db.
func (db *Database) GetAllTokens() ([]IPFSMetadata, error) {
	rows, err := db.Query("SELECT CID, Name, Description, ImageURL FROM metadata")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []IPFSMetadata
	for rows.Next() {
		var t IPFSMetadata
		if err := rows.Scan(&t.CID, &t.Name, &t.Description, &t.ImageURL); err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, nil
}

// GetTokenByCID retrieves a single token by its CID.
func (db *Database) GetTokenByCID(cid string) (*IPFSMetadata, error) {
	t := &IPFSMetadata{}
	err := db.QueryRow("SELECT CID, Name, Description, ImageURL FROM metadata WHERE CID = $1", cid).Scan(&t.CID, &t.Name, &t.Description, &t.ImageURL)
	if err != nil {
		return nil, err
	}
	return t, nil
}
