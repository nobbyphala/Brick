package rest_api

type VerifyDisbursementRequest struct {
	RecipientName          string `json:"recipient_name" validate:"gt=1,required"`
	RecipientAccountNumber string `json:"recipient_account_number" validate:"gte=1,numeric"`
	RecipientBankCode      string `json:"recipient_bank_code" validate:"gte=1"`
	Amount                 int64  `json:"amount"`
}

type DisburseRequest struct {
	RecipientName          string `json:"recipient_name" validate:"gt=1,required"`
	RecipientAccountNumber string `json:"recipient_account_number" validate:"gte=1,numeric"`
	RecipientBankCode      string `json:"recipient_bank_code" validate:"gte=1"`
	Amount                 int64  `json:"amount"`
}

type DisbursementResponse struct {
	Id                     string `json:"id"`
	RecipientName          string `json:"recipient_name"`
	RecipientAccountNumber string `json:"recipient_account_number"`
	RecipientBankCode      string `json:"recipient_bank_code"`
	Amount                 int64  `json:"amount"`
	Status                 string `json:"status"`
}

type BankTransferCallbackRequest struct {
	TransactionId string `json:"transaction_id" validate:"gte=1"`
	Status        string `json:"status" validate:"gte=1"`
}
