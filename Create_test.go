package directory

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		mode       os.FileMode
		shouldFail bool
	}{
		{
			name:       "Valid directory creation",
			path:       "/tmp/testdir",
			mode:       0755,
			shouldFail: false,
		},
		{
			name:       "Nested directory creation",
			path:       "/tmp/testdir/nested",
			mode:       0755,
			shouldFail: false,
		},
	}
	defer func() {
		if err := os.RemoveAll("/tmp/testdir"); err != nil {
			t.Fatal(err)
		}
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Create(tt.path, tt.mode)
			if (err != nil) != tt.shouldFail {
				t.Fatalf("Create() error = %v, shouldFail %v", err, tt.shouldFail)
			}

			// Check if directory exists if creation should not fail
			if !tt.shouldFail {
				if _, err := os.Stat(tt.path); os.IsNotExist(err) {
					t.Fatalf("expected directory %v to exist", tt.path)
				}
				// Clean up
				if err := os.RemoveAll(tt.path); err != nil {
					t.Fatalf("error removing file (%s): %v", tt.path, err)
				}
			}
		})
	}
}
