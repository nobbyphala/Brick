package database

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// for now the sqlx not wrapped

type PostgresSQLOpts struct {
	DB *sqlx.DB
}

type ConnectionOption struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type postgresSql struct {
	db *sqlx.DB
}

func NewPostgresDB(opts ConnectionOption) (*sqlx.DB, error) {
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

	return db, nil
}

func NewPostgresSqlClient(opts PostgresSQLOpts) *postgresSql {
	return &postgresSql{
		db: opts.DB,
	}
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

type PostgresSqlTxOpts struct {
	Tx *sqlx.Tx
}

func NewPostgresSqlTx(opts PostgresSqlTxOpts) *postgresSqlTx {
	return &postgresSqlTx{
		tx: opts.Tx,
	}
}

type postgresSqlTx struct {
	tx *sqlx.Tx
}

func (pgsx *postgresSqlTx) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgsx.tx.GetContext(ctx, dest, query, args...)
}

func (pgsx *postgresSqlTx) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return pgsx.tx.ExecContext(ctx, query, args...)
}

func (pgsx *postgresSqlTx) Query(ctx context.Context, query string, args ...interface{}) Row {
	return pgsx.tx.QueryRowContext(ctx, query, args...)
}
