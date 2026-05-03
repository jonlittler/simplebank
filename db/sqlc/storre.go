package db

import (
	"database/sql"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore defines all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new SQLStore (Store)
func NewStore(db *sql.DB) Store {
	return SQLStore{
		Queries: New(db),
		db:      db,
	}
}
