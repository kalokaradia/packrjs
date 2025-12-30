package config

import (
	"errors"
	"os"
	"path/filepath"
)

// Validate validates the configuration
func Validate(cfg *Config) error {
	if cfg.Input == "" {
		return errors.New("entry file is required")
	}

	switch cfg.SourceMap {
	case "none", "l", "in":
		return nil
	default:
		return errors.New("invalid sourcemap mode: use none, l, or in")
	}
}

// ValidateInputPath checks if the input path exists
func ValidateInputPath(input string) error {
	info, err := os.Stat(input)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("input path does not exist: " + input)
		}
		return err
	}
	if info.IsDir() {
		// For directory input, check if it's empty
		entries, err := os.ReadDir(input)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			return errors.New("input directory is empty: " + input)
		}
	}
	return nil
}

// ValidateOutputPath checks if the output parent directory exists
// Returns the parent directory path and an error if parent doesn't exist
func ValidateOutputPath(output string) (string, error) {
	dir := filepath.Dir(output)
	if dir == "." {
		// Output is in current directory, always valid
		return dir, nil
	}
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return dir, errors.New("output directory does not exist: " + dir)
		}
		return dir, err
	}
	if !info.IsDir() {
		return dir, errors.New("output path is not a directory: " + dir)
	}
	return dir, nil
}
