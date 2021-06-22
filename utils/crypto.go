package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// Base64 encodes a byte slice to a base64 string.
func Base64(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}

// GenRandomBytes generates crypto-secure random bytes.
func GenRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}
