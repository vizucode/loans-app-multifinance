package errorhandler

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validator kustom untuk email atau nomor telepon
func EmailOrPhoneValidator(fl validator.FieldLevel) bool {
	contact := fl.Field().String()

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	phoneRegex := regexp.MustCompile(`^(62|0)[0-9]{9,14}$`)

	return emailRegex.MatchString(contact) || phoneRegex.MatchString(contact)
}

func PasswordRegexValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Mengecek panjang password
	if len(password) < 8 {
		return false
	}

	// Mengecek apakah password mengandung setidaknya satu huruf besar
	var uppercasePattern = regexp.MustCompile(`[A-Z]`)
	if !uppercasePattern.MatchString(password) {
		return false
	}

	// Mengecek apakah password mengandung setidaknya satu angka
	var numberPattern = regexp.MustCompile(`[0-9]`)
	if !numberPattern.MatchString(password) {
		return false
	}

	// Mengecek apakah password mengandung setidaknya satu karakter khusus
	var specialCharPattern = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]{};':"\\|,.<>\/?]`)
	return specialCharPattern.MatchString(password)
}
