package rest_api

import (
	"github.com/gin-gonic/gin"
	"github.com/nobbyphala/Brick/domain/internal_error"
	"github.com/nobbyphala/Brick/external/validator"
	"net/http"
)

func SendErrorResponse(ctx *gin.Context, err error) {
	statusCode, exists := internal_error.ErrorStatusCodeMap[err.Error()]
	if !exists {
		statusCode = 500
	}

	resp := ErrorResponse{
		Message: err.Error(),
	}

	ctx.JSON(statusCode, resp)
}

func SendValidationErrorResponse(ctx *gin.Context, message string, errors []validator.ValidatorError) {
	var errorItems = make([]ValidationErrorItem, 0, len(errors))

	for _, err := range errors {
		errorItems = append(errorItems, ValidationErrorItem{
			Field: err.Field,
			Error: err.Error,
		})
	}

	response := ValidationErrorResponse{
		Message: message,
		Errors:  errorItems,
	}

	ctx.JSON(http.StatusBadRequest, response)
}
