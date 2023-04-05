package store

import (
	"database/sql"
	"time"
)

const (
	ctxTimeout = time.Second * 5
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
