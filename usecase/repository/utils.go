package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/nobbyphala/Brick/external/database"
)

type utils struct {
	db *sqlx.DB // Because sqlx not fully wrapped yet for now we will use sqlx directly
}

type UtilsOpts struct {
	DB *sqlx.DB
}

func NewRepositoryUtils(opts UtilsOpts) *utils {
	return &utils{
		db: opts.DB,
	}
}

func (ut utils) RunWithTransaction(ctx context.Context, handler func(Tx database.SQLDatabase) error) error {
	tx, err := ut.db.Beginx()
	if err != nil {
		return err
	}

	sqlDatabaseTx := database.NewPostgresSqlTx(database.PostgresSqlTxOpts{
		Tx: tx,
	})

	err = handler(sqlDatabaseTx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
