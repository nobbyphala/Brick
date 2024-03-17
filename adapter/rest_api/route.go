package rest_api

import "github.com/gin-gonic/gin"

type RouteController struct {
	DisbursementController *DisbursementController
}

func RegisterRouter(r *gin.Engine, ctrl RouteController) {
	r.POST("/disbursement/verify", ctrl.DisbursementController.VerifyDisbursement)
	r.POST("/disbursement", ctrl.DisbursementController.Disburse)
	r.PUT("/disbursement", ctrl.DisbursementController.HandleBankCallback)
}
