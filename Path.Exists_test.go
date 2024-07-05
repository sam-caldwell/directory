package directory

import (
	"testing"
)

// Unit test for Exists method
func TestPath_Exists(t *testing.T) {
	t.Run("Happy Path: /home exists", func(t *testing.T) {
		if path := Path("/home"); !path.Exists() {
			t.Fatalf("expected path exists, but function says it does not")
		}
	})
	t.Run("Happy Path: /home exists, expected: false", func(t *testing.T) {
		if path := Path("/does_not_exist"); path.Exists() {
			t.Fatalf("expected path does not exists, but function says it does")
		}
	})
}
