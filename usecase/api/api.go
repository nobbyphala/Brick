package api

import "context"

type Bank interface {
	VerifyAccount(ctx context.Context, account VerifyAccountRequest) (VerifyAccountResponse, error)
	TransferMoney(ctx context.Context, transfer TransferRequest) (TransferResponse, error)
}
