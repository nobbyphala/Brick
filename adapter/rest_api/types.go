package rest_api

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidationErrorItem struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationErrorResponse struct {
	Message string                `json:"message"`
	Errors  []ValidationErrorItem `json:"errors"`
}
