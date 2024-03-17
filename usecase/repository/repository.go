package repository

import (
	"context"
	"github.com/nobbyphala/Brick/usecase/repository/model"
)

type Disbursement interface {
	Insert(ctx context.Context, disbursement model.Disbursement) (string, error)
	UpdateById(ctx context.Context, id string, updatedData model.Disbursement) error
	GetByTransactionId(ctx context.Context, bankTransactionId string) (*model.Disbursement, error)
}
