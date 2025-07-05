package formatting

import "regexp"

func ReplaceSpecialCharacters(input string) string {
	// Regular expression untuk mendeteksi karakter khusus (selain huruf, angka, dan spasi)
	re := regexp.MustCompile(`[^a-zA-Z\s]`)
	return re.ReplaceAllString(input, "")
}
