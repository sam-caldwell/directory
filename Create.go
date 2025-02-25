package directory

import (
	"os"
)

// Create - Create a file-directory path with a given permission (mode)
//
//go:inline
func Create(path string, mode os.FileMode) (err error) {
	return os.MkdirAll(path, mode)
}
