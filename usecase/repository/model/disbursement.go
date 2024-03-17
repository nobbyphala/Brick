package model

import "time"

type Disbursement struct {
	Id                     string    `db:"id"`
	RecipientName          string    `db:"recipient_name"`
	RecipientAccountNumber string    `db:"recipient_account_number"`
	RecipientBankCode      string    `db:"recipient_bank_code"`
	TransferChannel        string    `db:"transfer_channel"`
	BankTransactionId      string    `db:"bank_transaction_id"`
	Amount                 int64     `db:"amount"`
	Status                 int       `db:"status"`
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
}
