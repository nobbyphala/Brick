package api

// some of these request and response struct is based on experienced in bank integration
type VerifyAccountStatus string
type TransferStatus string

const (
	AccountNotFoundStatus VerifyAccountStatus = "status: account not found"
	AccountBlockedStatus  VerifyAccountStatus = "status: account blocked"
	AccountVerifiedStatus VerifyAccountStatus = "status: account verified"
)

const (
	TransferStatusAccepted  TransferStatus = "ACCEPTED"
	TransferStatusRejected  TransferStatus = "REJECTED"
	TransferStatusFailed    TransferStatus = "FAILED"
	TransferStatusCompleted TransferStatus = "COMPLETED"
)

type VerifyAccountRequest struct {
	AccountHolderName   string `json:"account_holder_name"`
	AccountHolderNumber string `json:"account_holder_number"`
}

type VerifyAccountResponse struct {
	AccountHolderName   string              `json:"account_holder_name"`
	AccountHolderNumber string              `json:"account_holder_number"`
	AccountStatus       VerifyAccountStatus `json:"account_status"`
}

type TransferRequest struct {
	AccountHolderName   string `json:"account_holder_name"`
	AccountHolderNumber string `json:"account_holder_number"`
	DestinationBankCode string `json:"destination_bank_code"`
	Amount              int64  `json:"amount"`
}

type TransferResponse struct {
	TransactionId       string         `json:"transaction_id"`
	AccountHolderName   string         `json:"account_holder_name"`
	AccountHolderNumber string         `json:"account_holder_number"`
	DestinationBankCode string         `json:"destination_bank_code"`
	Amount              int            `json:"amount"`
	TransferStatus      TransferStatus `json:"transfer_status"`
}
