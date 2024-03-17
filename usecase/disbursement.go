package usecase

import (
	"context"
	"fmt"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/domain/internal_error"
	"github.com/nobbyphala/Brick/usecase/api"
	"github.com/nobbyphala/Brick/usecase/repository"
	"log"
)

type disbursementUsecase struct {
	bankApi                api.Bank
	disbursementRepository repository.Disbursement
}

type DisbursementDeps struct {
	BankApi                api.Bank
	DisbursementRepository repository.Disbursement
}

func NewDisbursement(deps DisbursementDeps) *disbursementUsecase {
	return &disbursementUsecase{
		bankApi:                deps.BankApi,
		disbursementRepository: deps.DisbursementRepository,
	}
}

func (disb disbursementUsecase) VerifyDisbursement(ctx context.Context, disbursement domain.Disbursement) error {
	verifyResponse, err := disb.bankApi.VerifyAccount(ctx, api.VerifyAccountRequest{
		AccountHolderName:   disbursement.RecipientName,
		AccountHolderNumber: disbursement.RecipientAccountNumber,
	})
	if err != nil {
		log.Println(err)
		return internal_error.ErrVerifyDisbursement
	}

	switch verifyResponse.AccountStatus {
	case api.AccountNotFoundStatus:
		return internal_error.ErrVerifyAccountNotFound
	case api.AccountBlockedStatus:
		return internal_error.ErrVerifyAccountBlocked
	case api.AccountVerifiedStatus:
		return nil
	default:
		return internal_error.ErrVerifyAccountNotFound
	}
}

func (disb disbursementUsecase) Disburse(ctx context.Context, disbursement domain.Disbursement) (domain.Disbursement, error) {
	err := disb.VerifyDisbursement(ctx, disbursement)
	if err != nil {
		log.Println(err)
		return domain.Disbursement{}, err
	}

	transferResponse, err := disb.bankApi.TransferMoney(ctx, api.TransferRequest{
		AccountHolderNumber: disbursement.RecipientAccountNumber,
		AccountHolderName:   disbursement.RecipientName,
		DestinationBankCode: disbursement.RecipientBankCode,
		Amount:              disbursement.Amount,
	})
	if err != nil {
		log.Println(err)
		return domain.Disbursement{}, internal_error.ErrDisburseBankError
	}

	insertedId, err := disb.disbursementRepository.Insert(ctx, domain.Disbursement{
		RecipientName:          disbursement.RecipientName,
		RecipientAccountNumber: disbursement.RecipientAccountNumber,
		RecipientBankCode:      disbursement.RecipientBankCode,
		BankTransactionId:      transferResponse.TransactionId,
		Amount:                 disbursement.Amount,
		Status:                 domain.DisbursementStatusPending,
	})
	if err != nil {
		log.Println(err)
		return domain.Disbursement{}, internal_error.ErrDisburseDisbursement
	}

	disbursement.Id = insertedId
	disbursement.BankTransactionId = transferResponse.TransactionId
	disbursement.Status = domain.DisbursementStatusPending

	return disbursement, nil
}

func (disb disbursementUsecase) ProcessBankCallback(ctx context.Context, bankCallback BankCallbackData) error {
	disbursement, err := disb.disbursementRepository.GetByTransactionId(ctx, bankCallback.TransactionId)
	if err != nil {
		log.Println(err)
		return internal_error.ErrHandleBankCallback
	}

	if disbursement == nil {
		log.Println("error disbursement not found")
		return internal_error.ErrDisbursementNotFound
	}

	if disbursement.Status != domain.DisbursementStatusPending {
		// invalid status need manual intervention
		err = internal_error.ErrDisbursementInvalidStatus
		log.Println(fmt.Sprintf("error status should be %d but got %d", domain.DisbursementStatusPending, disbursement.Status))
		return err
	}

	err = disb.disbursementRepository.UpdateById(ctx, disbursement.Id, domain.Disbursement{
		RecipientName:          disbursement.RecipientName,
		RecipientAccountNumber: disbursement.RecipientAccountNumber,
		RecipientBankCode:      disbursement.RecipientBankCode,
		BankTransactionId:      disbursement.BankTransactionId,
		Amount:                 disbursement.Amount,
		Status:                 disb.mapTransferStatusToDisbursementStatus(bankCallback.Status),
	})
	if err != nil {
		log.Println(err)
		return internal_error.ErrUpdateDisbursementStatus
	}

	return nil
}

func (disb disbursementUsecase) mapTransferStatusToDisbursementStatus(transferStatus api.TransferStatus) domain.DisbursementStatus {
	switch transferStatus {
	case api.TransferStatusCompleted:
		return domain.DisbursementStatusCompleted
	case api.TransferStatusAccepted:
		return domain.DisbursementStatusPending
	case api.TransferStatusFailed:
		return domain.DisbursementStatusFailed
	case api.TransferStatusRejected:
		return domain.DisbursementStatusRejected
	default:
		return domain.DisbursementStatusUnknown
	}
}
