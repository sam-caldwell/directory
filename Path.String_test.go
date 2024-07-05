package directory

import (
	"testing"
)

// Unit test for String method
func TestPath_String(t *testing.T) {
	tests := []struct {
		input    Path
		expected string
	}{
		{Path("home/user/docs"), "home/user/docs"},
		{Path(""), ""},
		{Path("/home/user/docs/"), "/home/user/docs/"},
		{Path("file.txt"), "file.txt"},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
