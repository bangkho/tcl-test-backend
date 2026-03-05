package helpers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single field validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface for ValidationErrors
func (v ValidationErrors) Error() string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Field+": "+err.Message)
	}
	return strings.Join(errs, ", ")
}

// validate is a singleton instance of validate
var validate = validator.New()

// SetTagName sets the struct tag name for validation
func SetTagName(tag string) {
	validate.SetTagName(tag)
}

// ValidateStruct validates a struct using go-playground/validator
// Returns ValidationErrors if validation fails, nil otherwise
func ValidateStruct(s interface{}) ValidationErrors {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var validationErrs ValidationErrors
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := strings.ToLower(err.Field())
		// Remove the first character if it's a capital letter to match struct field names
		if len(fieldName) > 0 {
			fieldName = strings.ToLower(string(fieldName[0])) + fieldName[1:]
		}

		validationErrs = append(validationErrs, ValidationError{
			Field:   fieldName,
			Message: getErrorMessage(err),
		})
	}

	return validationErrs
}

// ValidateVar validates a single variable
// Returns an error message if validation fails, empty string otherwise
func ValidateVar(field interface{}, tag string) string {
	err := validate.Var(field, tag)
	if err == nil {
		return ""
	}
	return getErrorMessage(err.(validator.ValidationErrors)[0])
}

// getErrorMessage returns a human-readable error message for a validation error
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short (minimum: " + err.Param() + ")"
	case "max":
		return "Value is too long (maximum: " + err.Param() + ")"
	case "len":
		return "Invalid length (expected: " + err.Param() + ")"
	case "numeric":
		return "Must be numeric"
	case "alphanum":
		return "Must contain only letters and numbers"
	case "uuid":
		return "Invalid UUID format"
	case "url":
		return "Invalid URL format"
	case "phone":
		return "Invalid phone number format"
	case "oneof":
		return "Value must be one of: " + err.Param()
	default:
		return "Invalid value for " + err.Field()
	}
}

// GetStructFieldName gets the actual field name from a struct pointer
func GetStructFieldName(s interface{}, tag string) string {
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("json") == tag {
			return field.Name
		}
	}
	return ""
}
