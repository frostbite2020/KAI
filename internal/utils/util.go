package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateToken() string {
	b := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
