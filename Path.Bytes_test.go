package directory

import (
	"testing"
)

// Unit test for Bytes method
func TestPath_Bytes(t *testing.T) {
	tests := []struct {
		input    Path
		expected []byte
	}{
		{Path("home/user/docs"), []byte("home/user/docs")},
		{Path(""), []byte("")},
		{Path("/home/user/docs/"), []byte("/home/user/docs/")},
		{Path("file.txt"), []byte("file.txt")},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := tt.input.Bytes()
			if string(result) != string(tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
