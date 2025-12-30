package utils

import (
	"fmt"
	"os"

	"github.com/kalokaradia/jspackr/src/cli"
)

// ConfirmOverwrite asks user for confirmation to overwrite a file
// If force or yes flag is set, skips confirmation
// Returns true if file should be overwritten, false otherwise
func ConfirmOverwrite(path string, force, yes, noConfirm bool) bool {
	// If noConfirm is set, auto-confirm everything
	if noConfirm {
		return true
	}

	// If force or yes is set, auto-overwrite existing files
	if force || yes {
		return true
	}

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}

	// Prompt user for confirmation
	message := fmt.Sprintf("Output file %s already exists. Overwrite?", path)
	return cli.Confirm(message, false)
}

