package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestGenerateSHA256Digest(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte("Hello, World!"), "315f5bdb76d078c43b8ac0064e4a016464a0b8d6a6d4b4b8b5b5b5b5b5b5b5b5"},
		{[]byte("GoLang"), "2c26b46b68ffc68ff99b453c1d30413413413413413413413413413413413413"},
		{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}

	for _, test := range tests {
		result := generateSHA256Digest(test.input)
		expectedHash := sha256.Sum256(test.input)
		expected := hex.EncodeToString(expectedHash[:])
		if result != expected {
			t.Errorf("generateSHA256Digest(%s) = %s; want %s", test.input, result, expected)
		}
	}
}
