// Package userdb contains user related CRUD functionality.
package userdb

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/slog"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log *slog.Logger
	db  *sqlx.DB
}

// NewStore constructs the api for data access.
func NewStore(log *slog.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}
