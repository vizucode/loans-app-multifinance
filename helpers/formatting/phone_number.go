package formatting

import "strings"

func PhoneTo62(phone string) string {
	// Trim whitespace
	phone = strings.TrimSpace(phone)

	// Jika nomor dimulai dengan "0", ubah menjadi "628"
	if strings.HasPrefix(phone, "0") {
		return "62" + phone[1:]
	}

	// Jika nomor sudah memiliki format internasional "+62", hilangkan "+"
	if strings.HasPrefix(phone, "+62") {
		return strings.Replace(phone, "+", "", 1)
	}

	// Jika nomor sudah dalam format internasional "62", kembalikan apa adanya
	if strings.HasPrefix(phone, "62") {
		return phone
	}

	// Default, kembalikan nomor apa adanya
	return phone
}
