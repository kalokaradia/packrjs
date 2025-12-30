package watcher

import (
	"crypto/sha256"
	"os"
)

// HashFile computes the SHA256 hash of a file
func HashFile(path string) ([32]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return [32]byte{}, err
	}
	return sha256.Sum256(data), nil
}
