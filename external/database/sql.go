package database

import (
	"context"
)

// wrap external depedencies and driver.
// for now we will define only the function needed for this project
type Row interface {
	Scan(dest ...any) error
	Err() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type SQLDatabase interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (Result, error)
	Query(ctx context.Context, query string, args ...interface{}) Row
}
