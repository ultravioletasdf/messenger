package main

import (
	"crypto/rand"
	"encoding/hex"
)

// Whether v is between min and max (inclusive)
func isBetween(v, min, max int) bool {
	return v >= min && v <= max
}

func generateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
