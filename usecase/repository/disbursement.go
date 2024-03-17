package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/external/database"
	"github.com/nobbyphala/Brick/usecase/repository/model"
)

type disbursementRepository struct {
	db database.SQLDatabase
}

type DisbursementDeps struct {
	DB database.SQLDatabase
}

func NewDisbursement(deps DisbursementDeps) *disbursementRepository {
	return &disbursementRepository{
		db: deps.DB,
	}
}

func (disb disbursementRepository) Insert(ctx context.Context, disbursement domain.Disbursement) (string, error) {
	var disbursementId string

	err := disb.db.Query(
		ctx,
		queryInsertDisbursement,
		disbursement.RecipientName,
		disbursement.RecipientAccountNumber,
		disbursement.RecipientBankCode,
		disbursement.BankTransactionId,
		disbursement.Amount,
		disbursement.Status,
	).Scan(&disbursementId)
	if err != nil {
		return "", err
	}

	return disbursementId, nil
}

func (disb disbursementRepository) UpdateById(ctx context.Context, id string, updatedData domain.Disbursement) error {
	_, err := disb.db.Exec(
		ctx,
		queryUpdateDisbursement,
		updatedData.RecipientName,
		updatedData.RecipientAccountNumber,
		updatedData.RecipientBankCode,
		updatedData.BankTransactionId,
		updatedData.Amount,
		updatedData.Status,
		id)

	return err
}

func (disb disbursementRepository) GetByTransactionId(ctx context.Context, bankTransactionId string) (*domain.Disbursement, error) {
	var res model.Disbursement

	err := disb.db.Get(ctx, &res, querySelectByBankTransactionId, bankTransactionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &domain.Disbursement{
		Id:                     res.Id,
		RecipientName:          res.RecipientName,
		RecipientAccountNumber: res.RecipientAccountNumber,
		RecipientBankCode:      res.RecipientBankCode,
		BankTransactionId:      res.BankTransactionId,
		Amount:                 res.Amount,
		Status:                 domain.DisbursementStatus(res.Status),
	}, nil
}
