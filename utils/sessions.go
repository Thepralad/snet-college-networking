package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSessionToken() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
