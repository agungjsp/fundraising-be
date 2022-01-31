package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta        `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type meta struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
}

func APIResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := response{
		Meta: meta,
		Data: data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
