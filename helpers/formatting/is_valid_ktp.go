package formatting

import "unicode"

func IsValidKTP(ktp string) bool {
	// Cek jika KTP kosong
	if ktp == "" {
		return false
	}

	// Cek jika panjang KTP lebih dari 26
	if len(ktp) > 26 {
		return false
	}

	// Cek jika KTP hanya berisi angka
	for _, r := range ktp {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	// Valid jika semua kondisi terpenuhi
	return true
}
