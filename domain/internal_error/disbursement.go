package internal_error

import "errors"

var (
	ErrVerifyDisbursement    = errors.New("error when trying verify disbursement")
	ErrVerifyAccountNotFound = errors.New("error bank account not found")
	ErrVerifyAccountBlocked  = errors.New("error bank account is blocked")

	ErrDisburseDisbursement = errors.New("error when try to disburse")
	ErrDisburseBankError    = errors.New("temporary bank network error")

	ErrHandleBankCallback       = errors.New("error when processing bank callback")
	ErrUpdateDisbursementStatus = errors.New("error when updated disbursement status")

	ErrDisbursementNotFound = errors.New("error disbursement not found")

	ErrDisbursementInvalidStatus = errors.New("error invalid disbursement status")
)
