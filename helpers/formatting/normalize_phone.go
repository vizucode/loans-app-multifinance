package formatting

import "strings"

func NormalizePhoneNumber(phone string) string {
	// Hapus tanda "+" jika ada di awal
	phone = strings.TrimPrefix(phone, "+")

	// // Jika nomor tidak berawal 62 atau 0, maka return string kosong
	// if !strings.HasPrefix(phone, "62") && !strings.HasPrefix(phone, "0") {
	// 	return ""
	// }

	// // Ganti awalan "0" dengan "62"
	// if strings.HasPrefix(phone, "0") {
	// 	return "62" + phone[1:]
	// }

	return phone
}
