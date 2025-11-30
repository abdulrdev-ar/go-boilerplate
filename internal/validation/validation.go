package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	if err == nil {
		return ""
	}

	var messages []string

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()

		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", field))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters", field, e.Param()))
		case "email":
			messages = append(messages, "Email is not valid")
		default:
			messages = append(messages, fmt.Sprintf("%s is invalid", field))
		}
	}

	return strings.Join(messages, ", ")
}
