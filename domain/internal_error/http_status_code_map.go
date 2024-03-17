package internal_error

import "net/http"

var ErrorStatusCodeMap = map[string]int{
	ErrInvalidRequest.Error():            http.StatusBadRequest,
	ErrVerifyDisbursement.Error():        http.StatusInternalServerError,
	ErrVerifyAccountNotFound.Error():     http.StatusOK,
	ErrVerifyAccountBlocked.Error():      http.StatusOK,
	ErrDisburseBankError.Error():         http.StatusInternalServerError,
	ErrDisburseDisbursement.Error():      http.StatusInternalServerError,
	ErrHandleBankCallback.Error():        http.StatusInternalServerError,
	ErrDisbursementNotFound.Error():      http.StatusNotFound,
	ErrDisbursementInvalidStatus.Error(): http.StatusNotFound,
	ErrUpdateDisbursementStatus.Error():  http.StatusInternalServerError,
}
