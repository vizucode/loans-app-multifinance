package formatting

import (
	"fmt"
	"time"
)

func GenerateTrxCode() string {
	now := time.Now()
	// Ambil detik sejak epoch dan ambil 10 digit
	trxCode := fmt.Sprintf("%02d%02d%02d%04d", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1e5)
	return trxCode
}
