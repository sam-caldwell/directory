package directory

import (
	"testing"
)

// Unit test for NewPath
func TestNewPath(t *testing.T) {
	t.Run("Happy path (test existing path, expecting exists)", func(t *testing.T) {
		path, err := NewPath[string]("/home", true)
		if err != nil {
			t.Fatalf("expected no error.  Got: %v", err)
		}
		if string(*path) != "/home" {
			t.Fatalf("path mismatch")
		}
	})

	t.Run("Happy path (test existing path, no expectation)", func(t *testing.T) {
		path, err := NewPath[string]("/home", false)
		if err != nil {
			t.Fatalf("expected no error.  Got: %v", err)
		}
		if string(*path) != "/home" {
			t.Fatalf("path mismatch")
		}
	})

	t.Run("Sad path (does not exist)(not expected)", func(t *testing.T) {
		path, err := NewPath[string]("/does_not_exist", false)
		if err != nil {
			t.Fatalf("expected no error.  Got: %v", err)
		}
		if string(*path) != "/does_not_exist" {
			t.Fatalf("path mismatch.  Got: %v", string(*path))
		}
	})

	t.Run("Sad path (does not exist)(expected)", func(t *testing.T) {
		path, err := NewPath[string]("/does_not_exist", true)
		if err == nil {
			t.Fatalf("expected error.")
		}
		if path != nil {
			t.Fatalf("path mismatch")
		}
	})
}
