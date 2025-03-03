package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateSHA256Digest(input []byte) string {
	hash := generateHash(input)
	return encodeToString(hash)
}

func generateHash(input []byte) []byte {
	hash := sha256.New()
	hash.Write(input)
	return hash.Sum(nil)
}

func encodeToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

func main() {
	fmt.Println("Copilot Testing")
}
