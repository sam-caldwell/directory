package directory

import (
	"testing"
)

// Unit test for Setp method
func TestPath_Setp(t *testing.T) {
	tests := []struct {
		initial  Path
		newValue string
		expected Path
	}{
		{Path("initial/path"), "new/path", Path("new/path")},
		{Path(""), "non/empty/path", Path("non/empty/path")},
		{Path("some/path"), "", Path("")},
	}

	for _, tt := range tests {
		t.Run(string(tt.initial), func(t *testing.T) {
			p := tt.initial
			p.Setp(&tt.newValue)
			if p != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, p)
			}
		})
	}
}
