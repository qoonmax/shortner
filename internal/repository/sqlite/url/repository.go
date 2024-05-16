package url

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"shortener/internal/repository"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveURL(urlToSave string, alias string) error {
	const fn = "repository.sqlite.SaveURL"

	stmt, err := r.db.Prepare("INSERT INTO urls (url, alias) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %v", fn, err)
	}

	_, err = stmt.Exec(urlToSave, alias)
	// TODO: Check if the alias already exists, return ErrURLAlreadyExists
	if err != nil {
		return fmt.Errorf("%s: failed to execute statement: %v", fn, err)
	}

	return nil
}

func (r *Repository) GetURL(alias string) (string, error) {
	const fn = "repository.sqlite.GetURL"

	var url string

	err := r.db.QueryRow("SELECT url FROM urls WHERE alias = ?", alias).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.ErrURLNotFound
		}

		return "", fmt.Errorf("%s: failed to get URL: %v", fn, err)
	}

	return url, nil
}
