package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(s interface{}) []ValidatorError
}

func NewValidator() *customValidator {
	return &customValidator{validate: validator.New()}
}

type customValidator struct {
	validate *validator.Validate
}

func (val customValidator) ValidateStruct(s interface{}) []ValidatorError {
	err := val.validate.Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)

		res := make([]ValidatorError, 0, len(validationErrors))

		for _, val := range validationErrors {
			res = append(res, ValidatorError{
				Field: val.Field(),
				Error: val.Error(),
			})
		}

		return res
	}

	return nil
}
