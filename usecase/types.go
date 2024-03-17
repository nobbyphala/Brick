package usecase

import "github.com/nobbyphala/Brick/usecase/api"

type BankCallbackData struct {
	TransactionId string
	Status        api.TransferStatus
}
