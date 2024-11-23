package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

type CheckHashFunc func(input, hashedInput string) bool

var CheckHash CheckHashFunc

func DefaultCheckHash(input, hashedInput string) bool {
	hash := Hash(input)
	return hash == hashedInput
}
