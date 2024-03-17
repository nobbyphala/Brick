package rest_api

import (
	"github.com/gin-gonic/gin"
	"github.com/nobbyphala/Brick/domain"
	"github.com/nobbyphala/Brick/domain/internal_error"
	"github.com/nobbyphala/Brick/external/validator"
	"github.com/nobbyphala/Brick/usecase"
	"github.com/nobbyphala/Brick/usecase/api"
	"net/http"
)

type DisbursementController struct {
	disbursementUsecase usecase.Disbursement
	validator           validator.Validator
}

type DisbursementControllerDeps struct {
	DisbursementUsecase usecase.Disbursement
}

func NewDisbursementController(deps DisbursementControllerDeps) *DisbursementController {
	return &DisbursementController{
		disbursementUsecase: deps.DisbursementUsecase,
		validator:           validator.NewValidator(),
	}
}

func (ctrl DisbursementController) VerifyDisbursement(ctx *gin.Context) {
	var requestBody VerifyDisbursementRequest

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		SendErrorResponse(ctx, internal_error.ErrInvalidRequest)
		return
	}

	validationErrors := ctrl.validator.ValidateStruct(requestBody)
	if validationErrors != nil {
		SendValidationErrorResponse(ctx, "invalid request", validationErrors)
		return
	}

	err = ctrl.disbursementUsecase.VerifyDisbursement(ctx.Request.Context(), domain.Disbursement{
		RecipientName:          requestBody.RecipientName,
		RecipientAccountNumber: requestBody.RecipientAccountNumber,
		RecipientBankCode:      requestBody.RecipientBankCode,
	})
	if err != nil {
		SendErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "disbursement successfully verified"})
}

func (ctrl DisbursementController) Disburse(ctx *gin.Context) {
	var requestBody DisburseRequest

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		SendErrorResponse(ctx, internal_error.ErrInvalidRequest)
		return
	}

	validationErrors := ctrl.validator.ValidateStruct(requestBody)
	if validationErrors != nil {
		SendValidationErrorResponse(ctx, "invalid request", validationErrors)
		return
	}

	disbursement, err := ctrl.disbursementUsecase.Disburse(ctx.Request.Context(), domain.Disbursement{
		RecipientName:          requestBody.RecipientName,
		RecipientAccountNumber: requestBody.RecipientAccountNumber,
		RecipientBankCode:      requestBody.RecipientBankCode,
		Amount:                 requestBody.Amount,
	})
	if err != nil {
		SendErrorResponse(ctx, err)
		return
	}

	response := DisbursementResponse{
		Id:                     disbursement.Id,
		RecipientName:          disbursement.RecipientName,
		RecipientAccountNumber: disbursement.RecipientAccountNumber,
		RecipientBankCode:      disbursement.RecipientBankCode,
		Amount:                 disbursement.Amount,
		Status:                 disbursement.Status.ToString(),
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl DisbursementController) HandleBankCallback(ctx *gin.Context) {
	var requestBody BankTransferCallbackRequest

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		SendErrorResponse(ctx, internal_error.ErrInvalidRequest)
		return
	}

	validationErrors := ctrl.validator.ValidateStruct(requestBody)
	if validationErrors != nil {
		SendValidationErrorResponse(ctx, "invalid request", validationErrors)
		return
	}

	err = ctrl.disbursementUsecase.ProcessBankCallback(ctx, usecase.BankCallbackData{
		TransactionId: requestBody.TransactionId,
		Status:        api.TransferStatus(requestBody.Status),
	})
	if err != nil {
		SendErrorResponse(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
