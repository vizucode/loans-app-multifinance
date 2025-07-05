package formatting

import (
	"crypto/rand"
	"math/big"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateETicketNumber() (string, error) {
	length := 6
	eTicket := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		eTicket[i] = charset[index.Int64()]
	}
	return string(eTicket), nil
}
