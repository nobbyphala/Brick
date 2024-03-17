package database

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresSQLOpts struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type postgresSql struct {
	db *sqlx.DB
}

func NewPostgresSql(opts PostgresSQLOpts) (*postgresSql, error) {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("user=%s dbname=%s host=%s port=%d password=%s sslmode=disable",
		opts.User,
		opts.Database,
		opts.Host,
		opts.Port,
		opts.Password,
	))
	if err != nil {
		return nil, err
	}

	return &postgresSql{
		db: db,
	}, nil
}

func (pgs *postgresSql) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgs.db.GetContext(ctx, dest, query, args...)
}

func (pgs *postgresSql) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return pgs.db.ExecContext(ctx, query, args...)
}

func (pgs *postgresSql) Query(ctx context.Context, query string, args ...interface{}) Row {
	return pgs.db.QueryRowContext(ctx, query, args...)
}
