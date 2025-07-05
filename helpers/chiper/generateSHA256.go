package chiper

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256(privateKey string, payload string) (resp string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	resp = string(hash.Sum(nil))
	return resp, nil
}

func GenerateSHA256Encoded(privateKey string, payload string) (resp string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	resp = hex.EncodeToString(hash.Sum(nil))
	return resp, nil
}
