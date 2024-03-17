package repository

import (
	"context"
	"github.com/nobbyphala/Brick/domain"
)

type Disbursement interface {
	Insert(ctx context.Context, disbursement domain.Disbursement) (string, error)
	UpdateById(ctx context.Context, id string, updatedData domain.Disbursement) error
	GetByTransactionId(ctx context.Context, bankTransactionId string) (*domain.Disbursement, error)
}
