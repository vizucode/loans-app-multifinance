package validatorize

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidatePhoneNumber is a custom validator for validating Indonesian phone numbers
func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	// Regex pattern for Indonesian phone numbers (min 10, max 15 characters)
	pattern := `^(\+628|628|08)[0-9]{7,12}$`

	// Match the phone number with regex
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}
