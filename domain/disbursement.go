package domain

type DisbursementStatus int

const (
	DisbursementStatusUnknown   DisbursementStatus = 0
	DisbursementStatusPending   DisbursementStatus = 1
	DisbursementStatusCompleted DisbursementStatus = 2
	DisbursementStatusFailed    DisbursementStatus = 3
	DisbursementStatusRejected  DisbursementStatus = 4
)

const (
	DisbursementStatusUnknownStr   = "UNKNOWN"
	DisbursementStatusPendingStr   = "PENDING"
	DisbursementStatusCompletedStr = "COMPLETED"
	DisbursementStatusFailedStr    = "FAILED"
	DisbursementStatusRejectedStr  = "REJECTED"
)

func (disb DisbursementStatus) ToString() string {
	switch disb {
	case DisbursementStatusUnknown:
		return DisbursementStatusUnknownStr
	case DisbursementStatusPending:
		return DisbursementStatusPendingStr
	case DisbursementStatusCompleted:
		return DisbursementStatusCompletedStr
	case DisbursementStatusFailed:
		return DisbursementStatusFailedStr
	case DisbursementStatusRejected:
		return DisbursementStatusRejectedStr
	default:
		return DisbursementStatusUnknownStr
	}
}

func (disb DisbursementStatus) ToInt() int {
	return int(disb)
}

type Disbursement struct {
	Id                     string
	RecipientName          string
	RecipientAccountNumber string
	RecipientBankCode      string
	TransferChannel        string // channel or bank used for doing the transfer
	Amount                 int64
	Status                 DisbursementStatus // status of the disbursement
}
