package sqlstore

import (
	"database/sql"
)

// Store ...
type Store struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
