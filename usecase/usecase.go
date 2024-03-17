package usecase

import (
	"context"
	"github.com/nobbyphala/Brick/domain"
)

type Disbursement interface {
	VerifyDisbursement(ctx context.Context, disbursement domain.Disbursement) error
	Disburse(ctx context.Context, disbursement domain.Disbursement) (domain.Disbursement, error)
	ProcessBankCallback(ctx context.Context, bankCallback BankCallbackData) error
}
