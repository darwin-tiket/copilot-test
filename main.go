package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateSHA256Digest(input []byte) string {
	hash := sha256.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	fmt.Println("Copilot Testing")
}
