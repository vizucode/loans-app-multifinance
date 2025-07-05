package formatting

import "strings"

// split name into 2 index ( firstname, lastname)
// ex: joko pramono rusdianto => [joko pramonom, rusdianto]
func SplitName(fullName string) (firstName, lastName string) {
	// Memecah input berdasarkan spasi
	words := strings.Fields(fullName)

	var firstname, lastname string

	if len(words) == 1 {
		// Jika hanya ada satu kata, set keduanya ke kata tersebut
		firstname = words[0]
		lastname = words[0]
	} else {
		// Jika ada lebih dari satu kata, firstname adalah gabungan kata pertama sampai kata terakhir kecuali kata terakhir
		firstname = strings.Join(words[:len(words)-1], " ") // Gabungkan semua kata kecuali kata terakhir
		lastname = words[len(words)-1]                      // Nama belakang adalah kata terakhir
	}

	return firstname, lastname
}
