package directory

import (
	"os"
	"testing"
)

func TestDelete(t *testing.T) {
	// Helper function to create a test directory with or without files
	createTestDir := func(path string, withFiles bool) error {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
		if withFiles {
			file, err := os.Create(path + "/testfile.txt")
			if err != nil {
				return err
			}
			if err := file.Close(); err != nil {
				return err
			}
		}
		return nil
	}

	tests := []struct {
		name           string
		path           string
		withFiles      bool
		deleteContents DeleteRule
		shouldFail     bool
	}{
		{
			name:           "Delete empty directory",
			path:           "emptydir",
			withFiles:      false,
			deleteContents: DeleteIfEmpty,
			shouldFail:     false,
		},
		{
			name:           "Delete non-empty directory with recursive delete",
			path:           "nonemptydir",
			withFiles:      true,
			deleteContents: RecursiveDelete,
			shouldFail:     false,
		},
		{
			name:           "Fail to delete non-empty directory without recursive delete",
			path:           "nonemptydir_no_recursive",
			withFiles:      true,
			deleteContents: DeleteIfEmpty,
			shouldFail:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createTestDir(tt.path, tt.withFiles); err != nil {
				t.Fatalf("Failed to create test directory: %v", err)
			}

			err := Delete(tt.path, tt.deleteContents)
			if (err != nil) != tt.shouldFail {
				t.Fatalf("Delete() error = %v, shouldFail %v", err, tt.shouldFail)
			}

			// Check if directory exists
			if _, err := os.Stat(tt.path); !os.IsNotExist(err) {
				if !tt.shouldFail {
					t.Fatalf("expected directory %v to be deleted", tt.path)
				} else {
					// Clean up the directory if the test expected it not to be deleted
					if err := os.RemoveAll(tt.path); err != nil {
						t.Fatalf("failed to clean up directory %v: %v", tt.path, err)
					}
				}
			}
		})
	}
}
