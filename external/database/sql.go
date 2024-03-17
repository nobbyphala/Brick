package database

import (
	"context"
	"database/sql"
)

// wrap external depedencies and driver.
// for now we will define only the function needed for this project
type SQLDatabase interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(ctx context.Context, query string, args ...interface{}) *sql.Row
}
