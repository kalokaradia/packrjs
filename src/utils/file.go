package utils

import (
	"crypto/sha256"
	"os"
	"path/filepath"
)

// DirNotEmpty checks if a directory is not empty
func DirNotEmpty(path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return len(entries) > 0, nil
}

// HashFile computes the SHA256 hash of a file
func HashFile(path string) ([32]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return [32]byte{}, err
	}
	return sha256.Sum256(data), nil
}

// CreateDir creates a directory recursively
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// GetOutputParent returns the parent directory of the output path
func GetOutputParent(output string) string {
	return filepath.Dir(output)
}

