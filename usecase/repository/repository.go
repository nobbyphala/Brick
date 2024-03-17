package repository

import (
	"context"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/external/database"
)

type Disbursement interface {
	WithTx(Tx database.SQLDatabase) Disbursement
	Insert(ctx context.Context, disbursement domain.Disbursement) (string, error)
	UpdateById(ctx context.Context, id string, updatedData domain.Disbursement) error
	GetByTransactionId(ctx context.Context, bankTransactionId string) (*domain.Disbursement, error)
}

type Utils interface {
	RunWithTransaction(ctx context.Context, handler func(Tx database.SQLDatabase) error) error
}
