package validator

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator interface {
	ValidateStruct(s interface{}) []ValidatorError
}

func NewValidator() *customValidator {
	en := en.New()
	uni := ut.New(en, en)

	translator, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, translator)

	return &customValidator{
		validate:   validate,
		translator: translator,
	}
}

type customValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func (custValidator customValidator) ValidateStruct(s interface{}) []ValidatorError {
	err := custValidator.validate.Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)

		res := make([]ValidatorError, 0, len(validationErrors))

		for _, val := range validationErrors {
			res = append(res, ValidatorError{
				Field: val.Field(),
				Error: val.Translate(custValidator.translator),
			})
		}

		return res
	}

	return nil
}
