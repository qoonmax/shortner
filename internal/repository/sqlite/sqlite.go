package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}
