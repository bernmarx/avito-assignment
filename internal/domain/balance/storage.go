package balance

import (
	sql "database/sql"
)

// Storage represents access to the database layer
type Storage struct {
	database
}

type database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Begin() (*sql.Tx, error)
}

// NewStorage creates new Storage
func NewStorage(db database) *Storage {
	return &Storage{database: db}
}
