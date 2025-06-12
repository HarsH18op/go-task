package my_validators

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidationMessage(fe validator.FieldError) string {
	param := fe.Param() // To access value of validation given
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "len":
		return "Length should be " + param
	case "min":
		return "Too short"
	case "max":
		return "Too long"
	case "gte":
		return "Must be greater than " + param
	case "lte":
		return "Must be less than or equal to " + param
	case "datetime":
		return "Must be a valid date (YYYY-MM-DD)"
	default:
		return "Invalid value"
	}
}
